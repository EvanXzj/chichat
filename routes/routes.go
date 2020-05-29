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
}
