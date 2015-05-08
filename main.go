package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"./config"
	"./logger"
	"./payload"
	"./validate"
)

func body2buffer(body io.ReadCloser) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	return buffer
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := config.GetConfig()
	body := body2buffer(r.Body)

	if !validate.ValidateRequest(body, r) {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid request")
		return
	}

	w.WriteHeader(200)

	payload := payload.Buffer2payload(body)
	logger.Log(c.Logfile, body.String())

	err := exec.Command(c.Execfile, payload.Repository.FullName).Run()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	c := config.GetConfig()
	http.HandleFunc("/", handler)
	http.ListenAndServe(c.BindAddress+":"+c.BindPort, nil)
}
