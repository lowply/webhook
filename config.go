package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

type Config struct {
	BindAddress string
	BindPort    string
	Execfile    string
	Logfile     string
	Key         string
}

const default_path = "/etc/webhook.hcl"

func NewConfig(path string) (*Config, error) {
	c := &Config{}

	if path == "" {
		path = default_path
	}

	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// Fail if config file is readable from other users
	if stat.Mode() != 0600 {
		return nil, errors.New("Permission bits of config file should be 0600.")
	}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	obj, err := hcl.Parse(string(d))
	if err != nil {
		return nil, err
	}

	err = hcl.DecodeObject(c, obj)
	if err != nil {
		return nil, err
	}

	if c.BindAddress == "" {
		return nil, errors.New("BindAddress is empty")
	}

	if c.BindPort == "" {
		return nil, errors.New("BindPort is empty")
	}

	if c.Execfile == "" {
		return nil, errors.New("Execfile is empty")
	}

	if c.Logfile == "" {
		return nil, errors.New("Logfile is empty")
	}

	if c.Key == "" {
		return nil, errors.New("Key is empty")
	}

	return c, nil
}
