package handlers

import (
	"net/http"

	"github.com/evanxzj/chitchat/models"
)

// Index home page
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		_, err := session(w, r)
		if err == nil {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "navbar", "index")
		}
	}
}
