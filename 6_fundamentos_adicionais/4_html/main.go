package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	u := usuario{ "Jose", "jose@email.com.br",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})
	fmt.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}