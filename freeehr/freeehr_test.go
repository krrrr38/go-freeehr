package freeehr

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()
	apiHandler := http.NewServeMux()
	apiHandler.Handle("/", mux)
	server := httptest.NewServer(apiHandler)
	client, _ = NewClient(http.DefaultClient)
	client.BaseURL, _ = url.Parse(server.URL)
	return client, mux, server.URL, server.Close
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}
