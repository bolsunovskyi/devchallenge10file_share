package router

import (
	"net/http"
	"file_share/handlers"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

var appRoutes = routes{
	route{
		Name:		"Front",
		Method:		"GET",
		Pattern:	"/",
		HandlerFunc:	handlers.Front,
	},
	route{
		Name:		"CreateUser",
		Method:		"POST",
		Pattern:	"/v1/user",
		HandlerFunc:	handlers.CreateUser,
	},
}
