package handlers

import (
	"fmt"
	"net/http"

	"github.com/evanxzj/chitchat/models"
)

// PostThread create a post
// POST /thread/post
func PostThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			danger("Cannot parse form")
			errorMessage(w, r, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			danger("Cannot get user from session")
			errorMessage(w, r, "Cannot get user from session")
		}

		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			danger("Cannot read thread")
			errorMessage(w, r, "Cannot read thread")
		}

		if _, err := user.CreatePost(thread, body); err != nil {
			danger("Cannot create post")
			errorMessage(w, r, err.Error())
		}
		url := fmt.Sprint("/thread/read?id=", thread.UUID)
		http.Redirect(w, r, url, 302)
	}
}
