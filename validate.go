package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func validateXhubsig(r *http.Request, sig string, key string) error {
	if r.Body == nil {
		return errors.New("Empty request body")
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// restore r.Body drained by ioutil.ReadAll.
	// https://medium.com/@xoen/golang-read-from-an-io-readwriter-without-loosing-its-content-2c6911805361
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	xhubsig := strings.Replace(sig, "sha1=", "", -1)
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(bodyBytes)

	if hex.EncodeToString(mac.Sum(nil)) != xhubsig {
		return errors.New("Hash calculation does not match to X-Hub-Signature")
	}

	return nil
}

func ValidateRequest(r *http.Request, key string) error {
	sig := r.Header.Get("X-Hub-Signature")

	if r.Method != "POST" {
		return errors.New("Request was not POST")
	}

	if sig == "" {
		return errors.New("X-Hub-Signature is empty")
	}

	err := validateXhubsig(r, sig, key)
	if err != nil {
		return err
	}

	return nil
}
