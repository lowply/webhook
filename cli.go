package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	ExitCodeOK    int    = 0
	ExitCodeError int    = 1 + iota
	Name          string = "webhook"
	Version       string = "0.0.2"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run() int {
	args := os.Args

	var (
		fl_config_path   string
		fl_generate_path string
		fl_version       bool
	)

	flags := flag.NewFlagSet("webhook", flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&fl_config_path, "c", "", "Path to the config file. Default value is /etc/webhook.hcl")
	flags.StringVar(&fl_generate_path, "g", "", "Path to generate config template file. No default value.")
	flags.BoolVar(&fl_version, "v", false, "Print version and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if fl_version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if fl_generate_path != "" {
		content := []byte(`# HCL file - https://github.com/hashicorp/hcl
bindaddress = "127.0.0.1"
bindport    = "4000"
execfile    = "/path/to/script.sh" # Make sure that the file is executable
logfile     = "/var/log/webhook.log"
key         = "" # GitHub webhook key. See https://developer.github.com/webhooks/securing/
`)
		err := ioutil.WriteFile(fl_generate_path, content, 0600)
		if err != nil {
			log.Fatalf("Error writing a file: ", err)
		}
		return ExitCodeOK
	}

	config, err := NewConfig(fl_config_path)
	if err != nil {
		log.Fatalf("Error reading config: ", err)
	}

	s := &Server{}
	server, err := s.NewServer(config)
	if err != nil {
		log.Fatalf("Error starting server: ", err)
	}

	l.InitLog(config.Logfile)
	server.Run()
	l.Close()

	return ExitCodeOK
}
