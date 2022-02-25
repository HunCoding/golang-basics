package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	httpObj := NewHttpInterface()

	http.Handle("/foo", httpObj)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
		case "PATCH":
		}
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type HttpStructImplementation struct {
}

func NewHttpInterface() http.Handler {
	return &HttpStructImplementation{}
}

func (h *HttpStructImplementation) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("123"))
	return
}
