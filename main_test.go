package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootExists(t *testing.T) {
	resp := get(t)
	if resp.StatusCode == 404 {
		t.Fatalf("want 404, got %d", resp.StatusCode)
	}
}

func TestRootReturnsOk(t *testing.T) {
	resp := get(t)
	if resp.StatusCode != 200 {
		t.Fatalf("want 200, got %d", resp.StatusCode)
	}
}

func TestRootContentTypeHeader(t *testing.T) {
	resp := get(t)
	v := resp.Header.Get("Content-Type")
	if v != "text/html" {
		t.Fatalf("want application/json, got %s", v)
	}
}

func TestRootContentLength(t *testing.T) {
	resp := get(t)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := string(b)
	if c != "Hello World!" {
		t.Fatalf("want Hello World!, got %s", c)
	}
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
