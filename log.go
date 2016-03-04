package main

import (
	"log"
	"os"

	"github.com/google/logger"
)

type Log struct {
	file *os.File
}

func (l *Log) InitLog(path string) {
	var err error
	l.file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	l.Init()
}

func (l *Log) Init() {
	logger.Init("WebhookLogger", false, true, l.file)
}

func (l *Log) Close() {
	l.file.Close()
}
