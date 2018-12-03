package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootReturnsOk(t *testing.T) {
	resp := get(t, nil)
	assert(t, 200, resp.StatusCode)
}

func TestRootContentTypeHeader(t *testing.T) {
	resp := get(t, nil)
	v := resp.Header.Get("Content-Type")
	assert(t,"application/json", v)
}

func TestRootParameters(t *testing.T) {
	testCases := []struct{
		given Parameters
		want string
	}{
		{given: Parameters{}, want: "Hello, 世界!"},
		{given: Parameters{greeting: "Sup"}, want: "Sup, 世界!"},
		{given: Parameters{name: "World"}, want: "Hello, World!"},
		{given: Parameters{greeting: "Sup", name: "World"}, want: "Sup, World!"},
	}
	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			m := getMessageForParameters(t, &tc.given)
			assert(t, tc.want, m)
		})
	}
}

func TestEasterEgg(t *testing.T) {
	parameters := []Parameters{
		{greeting: "Sup", name: "Son"},
		{greeting: "Sup", name: "son"},
		{greeting: "sup", name: "Son"},
		{greeting: "sup", name: "son"},
	}
	for _, p := range parameters {
		t.Run(fmt.Sprintf("%s, %s!", p.greeting, p.name), func(t *testing.T) {
			m := getMessageForParameters(t, &p)
			assert(t, "¯\\_(ツ)_/¯", m)
		})
	}
}

func getMessageForParameters(t *testing.T, p *Parameters) string {
	r := get(t, p)
	decoder := json.NewDecoder(r.Body)
	var b Body
	if err := decoder.Decode(&b); err != nil {
		log.Fatal(err)
	}
	return b.Message
}

func get(t *testing.T, p *Parameters) *http.Response {
	ts := httptest.NewServer(http.HandlerFunc(handle))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	if p != nil {
		q.Add("greeting", p.greeting)
		q.Add("name", p.name)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := ts.Client().Do(req)
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

type Parameters struct {
	greeting string
	name     string
}