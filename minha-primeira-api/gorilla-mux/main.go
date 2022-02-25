package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(
		rw http.ResponseWriter,
		req *http.Request) {

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("123"))
		return
	}).Methods("GET", "POST")

	http.ListenAndServe(":8081", r)
}
