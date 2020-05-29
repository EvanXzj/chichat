package handlers

import (
	"fmt"
	"net/http"

	"github.com/evanxzj/chitchat/models"
)

// Login login handler
// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	t := parseTemplateFiles("auth.layout", "navbar", "login")
	t.Execute(w, nil)
}

// Signup signup page handler
// GET /signup
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}

// SignupAccount post signup
// POST /signup
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Cannot parse form")
	}

	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		fmt.Println("Cannot create user")
	}

	http.Redirect(w, r, "/login", 302)
}

// Authenticate post login
// POST /authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("can not parse form111")
	}
	e := r.PostFormValue("email")
	user, err := models.UserByEmail(e)
	if err != nil {
		fmt.Println("can not find user")
	}

	if user.Password == models.Sha1(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println(err)
			fmt.Println("can not create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// Logout user logout handler
// GET logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		fmt.Println("Failed to get cookie")
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.Delete()
	}
	http.Redirect(w, r, "/", 302)
}
