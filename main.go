package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/lowply/webhook/config"
	"github.com/lowply/webhook/logger"
	"github.com/lowply/webhook/payload"
	"github.com/lowply/webhook/validate"
)

func body2buffer(body io.ReadCloser) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	return buffer
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := config.GetConfig()
	body := body2buffer(r.Body)

	err := validate.ValidateRequest(body, r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid request\n")
		logger.Log("Invalid request from " + r.Header.Get("X-Forwarded-For") + ": " + err.Error())
		return
	}

	w.WriteHeader(200)
	logger.Log("Valid request from " + r.Header.Get("X-Forwarded-For") + ": " + body.String())

	payload := payload.Buffer2payload(body)

	err = exec.Command(c.Execfile, payload.Repository.FullName).Run()
	if err != nil {
		logger.Log("Error running script: " + c.Execfile)
	}

}

func main() {
	c := config.GetConfig()
	logger.Log("webhook started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(c.BindAddress+":"+c.BindPort, nil)
}
