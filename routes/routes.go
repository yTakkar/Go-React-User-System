package routes

import (
	"html/template"
	"log"
	"net/http"
)

// Page type
type Page struct {
	Title    string
	userID   interface{}
	username interface{}
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		log.Fatal(err)
	}
}
