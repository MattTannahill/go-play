package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootReturnsOk(t *testing.T) {
	resp := get(t, "")
	assert(t, 200, resp.StatusCode)
}

func TestRootContentTypeHeader(t *testing.T) {
	resp := get(t, "")
	v := resp.Header.Get("Content-Type")
	assert(t,"text/html", v)
}

func TestRootContentLength(t *testing.T) {
	c := getAndDeserialize(t, "")
	assert(t, "Hello, World!", c)
}

func TestRootNameParameter(t *testing.T) {
	c := getAndDeserialize(t, "Matt")
	assert(t, "Hello, Matt!", c)
}

func getAndDeserialize(t *testing.T, name string) string {
	resp := get(t, name)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	c := string(b)
	return c
}

func get(t *testing.T, name string) *http.Response {
	ts := httptest.NewServer(http.HandlerFunc(handle))
	defer ts.Close()

	url := ts.URL + "?"
	if name != "" {
		url += "name=" + name
	}

	resp, err := http.Get(url)
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