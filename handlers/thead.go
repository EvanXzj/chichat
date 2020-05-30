package handlers

import (
	"net/http"

	"github.com/evanxzj/chitchat/models"
)

// NewThread new thread page
// GET /threads/new
func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "auth.navbar", "new.thread")
	}
}

// CreateThread create new thread
// POST /thread/create
func CreateThread(w http.ResponseWriter, r *http.Request) {
	session, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger("can not parse form")
			errorMessage(w, r, "Cannot  parse form")
			return
		}

		user, err := session.User()
		if err != nil {
			danger("failed get user info from session")
			errorMessage(w, r, "failed get user info from session")
			return
		}
		topic := r.PostFormValue("topic")
		if _, err = user.CreateThread(topic); err != nil {
			danger("Cannot create thread")
			errorMessage(w, r, "failed get user info from session")
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}

// ReadThread read thread by uuid
// GET /thread/read
func ReadThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		danger("can not read the thread")
		errorMessage(w, r, "Cannot read thread")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(w, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
