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
			fmt.Println("Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			fmt.Println("Cannot get user from session")
		}

		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			fmt.Println("Cannot read thread")
		}

		if _, err := user.CreatePost(thread, body); err != nil {
			fmt.Println(err)
			fmt.Println("Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", thread.UUID)
		http.Redirect(w, r, url, 302)
	}
}
