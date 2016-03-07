package main

import (
	"net/http"
	"os/exec"

	"github.com/google/logger"
)

type Server struct {
	bindAddress string
	bindPort    string
	execFile    string
	logFile     string
	key         string
}

func (self *Server) handler(w http.ResponseWriter, r *http.Request) {

	err := ValidateRequest(r, self.key)
	if err != nil {
		w.WriteHeader(400)
		logger.Errorf("Invalid request: %v", err)
		return
	}

	w.WriteHeader(200)
	logger.Info("Valid request from: " + r.Header.Get("User-Agent"))
	logger.Info("Delivery ID: " + r.Header.Get("X-GitHub-Delivery"))

	payload, err := NewPayload(r.Body)
	if err != nil {
		logger.Errorf("Error creating payload: %v", err)
	}

	err = exec.Command(self.execFile, payload.Repository.FullName).Run()
	if err != nil {
		logger.Errorf("Error running script: %v", err)
	} else {
		logger.Info("Ran script: " + self.execFile)
	}
}

func NewServer(config *Config) (*Server, error) {
	s := &Server{}
	s.bindAddress = config.BindAddress
	s.bindPort = config.BindPort
	s.execFile = config.Execfile
	s.logFile = config.Logfile
	s.key = config.Key
	return s, nil
}

func (self *Server) Run() {
	logger.Info("Webhook server started")
	http.HandleFunc("/", self.handler)
	err := http.ListenAndServe(self.bindAddress+":"+self.bindPort, nil)
	if err != nil {
		logger.Fatalf("Error running server: %v", err)
	}
}
