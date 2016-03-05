# webhook
Receive webhook from Github and run a script

## Description

Simple http server to accept GitHub's webhook and run any script you like. It only accepts POST with valid `X-Hub-Signature` http header. Then it passes `repository.full_name` (repo's username/reponame string) as an argument to a script. Script can be any language you like.

## Example

For example you can use it to automate re-building a docker image and restarting a container when you push your code to the repository.

Config example:

```bash
# /etc/webhook.hcl

bindaddress	= "127.0.0.1"
bindport	= "4000"
execfile	= "/home/docker/pull.sh" # Make sure that the file is executable
logfile		= "/var/log/webhook.log"
key         = "XXXXXXXXXXXXXXXXXXXX" # GitHub webhook key. See https://developer.github.com/webhooks/securing/
```

Example script for docker:
```bash
#!/bin/bash

# /home/docker/pull.sh

cd $GOPATH/src/github.com/${1}
git pull origin master
docker build -t .
docker stop
docker rm
docker start
```

## Installation
```bash
$ go get github.com/lowply/webhook
```

## Start webhook server
```bash
$ webhook -g /etc/webhook.hcl # generate config template
$ webhook # defalut config path is /etc/webhook.hcl
```

## Specify config path
```bash
$ webhook -c /path/to/webhook.hcl # if you need to specify path of config file
```

## Daemonize
I recommend supervisor to daemonize webhook and setting up a reverse proxy with nginx.

Supervisor example:
```
# /etc/supervisord.d/webhook.conf
[program:webhook]
command=[GOPATH]/bin/webhook # change GOPATH to yours
directory=/tmp
user=root
stdout_logfile=/var/log/supervisor/webhook.stdout.log
stderr_logfile=/var/log/supervisor/webhook.stderr.log
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

## Author
[Sho Mizutani](https://github.com/lowply)
