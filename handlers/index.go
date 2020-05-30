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

// Error error page
// GET /error?msg=xxxx
func Error(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}

// PrintVersion version print
func PrintVersion(w http.ResponseWriter, r *http.Request) {
	version := []byte(Version())
	w.Write(version)
}
