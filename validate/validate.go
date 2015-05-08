package validate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"

	"github.com/lowply/webhook/config"
)

func validateXhubsig(b *bytes.Buffer, sig string) error {
	c := config.GetConfig()

	xhubsig := strings.Replace(sig, "sha1=", "", -1)
	key := c.Key
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(b.Bytes())

	if hex.EncodeToString(mac.Sum(nil)) != xhubsig {
		return errors.New("")
	}

	return nil
}

func ValidateRequest(b *bytes.Buffer, r *http.Request) error {
	sig := r.Header.Get("X-Hub-Signature")

	if r.Method != "POST" {
		return errors.New("Request was not POST")
	}

	if sig == "" {
		return errors.New("X-Hub-Signature is empty")
	}

	err := validateXhubsig(b, sig)
	if err != nil {
		return errors.New("Invalid X-Hub-Signature")
	}

	return nil
}
