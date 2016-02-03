package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func body2buffer(body io.ReadCloser) *bytes.Buffer {
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(body)
	return buffer
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := GetConfig()
	body := body2buffer(r.Body)

	err := ValidateRequest(body, r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Invalid request\n")
		Log("Invalid request from " + r.Header.Get("X-Forwarded-For") + ": " + err.Error())
		return
	}

	w.WriteHeader(200)
	Log("Valid request from " + r.Header.Get("X-Forwarded-For") + ": " + body.String())

	payload := Buffer2payload(body)

	err = exec.Command(c.Execfile, payload.Repository.FullName).Run()
	if err != nil {
		Log("Error running script: " + c.Execfile)
	}

}

func main() {
	c := GetConfig()
	Log("webhook started")
	http.HandleFunc("/", handler)
	http.ListenAndServe(c.BindAddress+":"+c.BindPort, nil)
}
