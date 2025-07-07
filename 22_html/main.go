package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		usuario := usuario{
			Nome:  "João",
			Email: "joao@example.com",
		}
		templates.ExecuteTemplate(w, "home.html", usuario) // Renderiza o template "home.html" passando o objeto usuario
	})

	log.Fatal(http.ListenAndServe(":5000", nil)) // Inicia um servidor HTTP na porta 5000

}
