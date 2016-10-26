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
	route{
		Name:		"LoginUser",
		Method:		"POST",
		Pattern:	"/v1/token",
		HandlerFunc:	handlers.LoginUser,
	},
	route{
		Name:		"UploadFile",
		Method:		"POST",
		Pattern:	"/v1/file/{fileName:[0-9a-zA-Z._]+}",
		HandlerFunc:	handlers.UploadFile,
	},
	route{
		Name:		"ListFiles",
		Method:		"GET",
		Pattern:	"/v1/files/{parent}",
		HandlerFunc:	handlers.ListFiles,
	},
	route{
		Name:		"ListFilesNoParent",
		Method:		"GET",
		Pattern:	"/v1/files",
		HandlerFunc:	handlers.ListFiles,
	},
}
