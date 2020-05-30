package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/evanxzj/chitchat/models"
)

var logger *log.Logger

func init() {
	file, err := os.OpenFile("./logs/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func info(args ...interface{}) {
	logger.SetPrefix("INFO")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING")
	logger.Println(args...)
}

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

func errorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/error?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}
