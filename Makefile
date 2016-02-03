
all: run

run:
	GO15VENDOREXPERIMENT=1 go run *.go

build:
	GO15VENDOREXPERIMENT=1 go build -o bin/webhook

install:
	go install github.com/lowply/webhook

install-linux:
	GOOS=linux GOARCH=amd64 go install github.com/lowply/webhook
