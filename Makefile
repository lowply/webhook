default: test

run:
	go run *.go

build:
	go build -o bin/webhook

install:
	go install github.com/lowply/webhook

install-linux:
	GOOS=linux GOARCH=amd64 go install github.com/lowply/webhook
	
vendor:
	govend -v
