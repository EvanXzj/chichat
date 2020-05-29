package handlers

import (
	"html/template"
	"net/http"

	"github.com/evanxzj/chitchat/models"
)

// Index home page
func Index(w http.ResponseWriter, r *http.Request) {
	files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := models.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
