package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
        port := os.Getenv("PORT")
	if port == "" {
		port = "8080";
	}

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":" + port,  nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	greeting, name := getParameterOrFallback(r, "greeting", "Hello"), getParameterOrFallback(r, "name", "世界")
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
	if strings.EqualFold(greeting, "sup") && strings.EqualFold(name, "son") {
		return "¯\\_(ツ)_/¯"
	}
	return fmt.Sprintf("%s, %s!", greeting, name)
}

type Body struct {
	Message string
}
