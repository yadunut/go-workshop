package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	addr = ":8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintf(w, "Hello World!")
		return
	}
	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	http.HandleFunc("/greet/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
