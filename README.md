# webhook
Receive webhook from Github and run a script

## Description

Simple http server to accept GitHub's webhook and run any script you like. It only accepts POST with valid `X-Hub-Signature` http header. Then it passes `repository.full_name` (repo's username/reponame string) as an argument to a script. Script can be any language you like.

## Example

For example you can use it to automate re-building a docker image and restarting a container when you push your code to the repository.

Config example:

```bash
# /etc/webhook.hcl

bindaddress     = "0.0.0.0" # be careful, this is just an example.
bindport        = "4000"
execfile        = "/home/pull.sh"
logfile         = "/var/log/webhook.log"
key             = "XXXXXXXXXXXXXXXXXXXX" # GitHub webhook key
```

Script example:
```bash
#!/bin/bash

# /home/pull.sh

cd $GOPATH/src/github.com/${1}
git pull origin master
docker build -t .
docker stop
docker rm
docker start
```

With this config and script, your docker build and restart will be automated.

## Installation

Just clone, build and run in the background. You will need to install git and go beforehand.

```bash
$ git clone https://github.com/lowply/webhook.git # clone it
$ cd webhook
$ go build main.go # build it
$ cp webhook.hcl.tmpl /etc/webhook.hcl # copy it
$ chmod 600 /etc/webhook.hcl # permission should be 600
$ vim /etc/webhook.hcl # update
$ ./main & # run
```

However I recommend using supervisor to daemonize it, and setting up reverse proxy with nginx.

Supervisor example:
```
# /etc/supervisord.d/webhook.conf
[program:webhook]
command=/path/to/webhook/main
directory=/tmp
user=root
stdout_logfile=/var/log/supervisor/webhook.stdout.log
stderr_logfile=/var/log/supervisor/webhook.stderr.log
```

Start:
```bash
$ supervisorctl status
$ supervisorctl reread
webhook: available
$ supervisorctl add webhook
webhook: added process group
$ supervisorctl status
webhook                          RUNNING   pid 27606, uptime 0:00:02
```
Webhook example:
```bash
# /etc/webhook.hcl

bindaddress     = "127.0.0.1"
bindport        = "4000"
execfile        = "/path/to/script.sh"
logfile         = "/var/log/webhook.log"
key             = "XXXXXXXXXXXXXXXXXXXX" # GitHub webhook key
```

Nginx example:
```
# /etc/nginx/conf.d/hook.example.com.conf
upstream hook.example.com {
        server localhost:4000;
}

server {
    listen       80;
    server_name  hook.example.com;
    access_log  /var/log/nginx/hook.example.com.access.log  main;
    location / {
        proxy_pass http://hook.example.com;
    }
}
```

## TODO

- Write test
- Support for multiple repositories
- Binary distribution instead of repository cloning


## Author
[Sho Mizutani](https://github.com/lowply)
