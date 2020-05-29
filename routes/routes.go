package routes

import (
	"net/http"

	"github.com/evanxzj/chitchat/handlers"
)

// WebRoute a single route struct
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signupPage",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	{
		"loginPage",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
}
