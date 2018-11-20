package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootReturnsOk(t *testing.T) {
	resp := get(t)
	assert(t, 200, resp.StatusCode)
}

func TestRootContentTypeHeader(t *testing.T) {
	resp := get(t)
	v := resp.Header.Get("Content-Type")
	assert(t,"text/html", v)
}

func TestRootContentLength(t *testing.T) {
	resp := get(t)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := string(b)
	assert(t,"Hello World!", c)
}

func get(t *testing.T) *http.Response {
	ts := httptest.NewServer(http.HandlerFunc(handle))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}

func assert(t *testing.T, want interface{}, got interface{}) {
	if want != got {
		t.Fatalf("want %v, got %v", want, got)
	}
}