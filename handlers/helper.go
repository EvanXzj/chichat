package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/evanxzj/chitchat/models"
)

// check user logged or not with Cookie
func session(w http.ResponseWriter, r *http.Request) (session models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		session = models.Session{UUID: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}

	return
}

// parse template with file names and return return template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	t = template.Must(t.ParseFiles(files...))
	return
}

// generate html and response to client
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	t := template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", data)
}

// Version return current app version number
func Version() string {
	return "0.0.1"
}
