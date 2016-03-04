package main

import (
	"flag"
	"fmt"
	"io"
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

	flags.StringVar(&fl_config_path, "c", "", "[c]onfig: this is fl_config_path")
	flags.StringVar(&fl_generate_path, "g", "", "[g]enerate: Generate config file to a path")
	flags.BoolVar(&fl_version, "v", false, "Print version and quit.")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	if fl_version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if fl_generate_path != "" {
		fmt.Println("do generate")
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
