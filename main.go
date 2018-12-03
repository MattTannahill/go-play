package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080",  nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	greeting := getParameterOrFallback(r, "greeting", "Hello")
	name := getParameterOrFallback(r, "name", "World")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Body{
		Message: getMessage(greeting, name),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getParameterOrFallback(r *http.Request, key, fallback string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		value = fallback
	}
	return value
}

func getMessage(greeting, name string) string {
	if strings.ToLower(greeting) == "sup" && strings.ToLower(name) == "son" {
		return "¯\\_(ツ)_/¯"
	}
	return fmt.Sprintf("%s, %s!", greeting, name)
}

type Body struct {
	Message string
}