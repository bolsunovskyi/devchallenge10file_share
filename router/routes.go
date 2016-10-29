package router

import (
	"net/http"
	"file_share/handlers"
	"file_share/middleware"
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
		HandlerFunc:	middleware.Auth(handlers.UploadFile),
	},
	route{
		Name:		"SearchFiles",
		Method:		"GET",
		Pattern:	"/v1/files/search/{keyword}",
		HandlerFunc:	middleware.Auth(handlers.SearchFiles),
	},
	route{
		Name:		"ListFiles",
		Method:		"GET",
		Pattern:	"/v1/files/{parent}",
		HandlerFunc:	middleware.Auth(handlers.ListFiles),
	},
	route{
		Name:		"ListFilesNoParent",
		Method:		"GET",
		Pattern:	"/v1/files",
		HandlerFunc:	middleware.Auth(handlers.ListFiles),
	},
	route{
		Name:		"MoveFile",
		Method:		"PATCH",
		Pattern:	"/v1/file/{fileID}",
		HandlerFunc:	middleware.Auth(handlers.MoveFile),
	},
	route{
		Name:		"DeleteFile",
		Method:		"DELETE",
		Pattern:	"/v1/file/{fileID}",
		HandlerFunc:	middleware.Auth(handlers.DeleteFile),
	},
	route{
		Name:		"RenameFile",
		Method:		"PUT",
		Pattern:	"/v1/file/{fileID}",
		HandlerFunc:	middleware.Auth(handlers.RenameFile),
	},
	route{
		Name:		"DownloadFile",
		Method:		"GET",
		Pattern:	"/v1/file/{fileID}",
		HandlerFunc:	middleware.Auth(handlers.DownloadFile),
	},
}
