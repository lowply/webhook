package config

import (
	"io/ioutil"
	"log"
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

func GetConfig() Config {
	var c Config
	path := "/etc/webhook.hcl"

	// Fail if config does not exist
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalf("Error reading %s: %s", path, err)
	}

	// Fail if config file is readable from other users
	if stat.Mode() != 0600 {
		log.Fatalf("Permission bits of config file %s should be 0600.", path)
	}

	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading %s: %s", path, err)
	}

	obj, err := hcl.Parse(string(d))
	if err != nil {
		log.Fatalf("Error parsing %s: %s", obj, err)
	}

	hcl.DecodeObject(&c, obj)
	if err != nil {
		log.Fatalf("Error decoding %s: %s", path, err)
	}

	return c
}
