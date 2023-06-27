package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ola servidor HTTP"))
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Usuario"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}