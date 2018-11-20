package main

import (
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
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s, %s!", greeting, name)
}

func getParameterOrFallback(r *http.Request, key, fallback string) string {
	v := r.URL.Query().Get(key)
	if v == "" {
		v = fallback
	}
	return v
}