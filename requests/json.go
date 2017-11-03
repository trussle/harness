package requests

import (
	"bytes"
	"net/http"
	"testing"
)

// Get sends a request to the URL with the ContentType of application/json and
// returns the raw bytes which can be decoded
func Get(t *testing.T, url string) []byte {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")

	return Request(t, Build(
		WithHeaders(header),
		WithMethod("GET"),
		WithURL(url),
		WithStatusCode(200),
	))
}

// Post sends a request to the URL with the ContentType of application/json and
// returns the raw bytes which can be decoded
func Post(t *testing.T, url string, body []byte) []byte {
	header := make(http.Header)
	header.Set("Content-Type", "application/json")

	return Request(t, Build(
		WithHeaders(header),
		WithMethod("GET"),
		WithURL(url),
		WithBody(bytes.NewReader(body)),
		WithStatusCode(200),
	))
}
