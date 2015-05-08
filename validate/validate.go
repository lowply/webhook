package validate

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strings"

	"../config"
)

func ValidateRequest(b *bytes.Buffer, r *http.Request) bool {
	c := config.GetConfig()

	if r.Method != "POST" {
		return false
	}

	if r.Header.Get("X-Hub-Signature") == "" {
		return false
	}

	xhubsig := strings.Replace(r.Header.Get("X-Hub-Signature"), "sha1=", "", -1)
	key := c.Key
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write(b.Bytes())

	// Check Signature
	if hex.EncodeToString(mac.Sum(nil)) != xhubsig {
		return false
	}

	return true
}
