package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
		Message: fmt.Sprintf("%s, %s!", greeting, name),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getParameterOrFallback(r *http.Request, key, fallback string) string {
	v := r.URL.Query().Get(key)
	if v == "" {
		v = fallback
	}
	return v
}

type Body struct {
	Message string
}