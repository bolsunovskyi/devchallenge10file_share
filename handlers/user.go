package handlers

import (
	"net/http"
	"file_share/repository/user"
	"file_share/jwt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	appUser, err := user.CreateUser(
		r.FormValue("first_name"),
		r.FormValue("last_name"),
		r.FormValue("email"),
		r.FormValue("password"))

	if err != nil {
		sendError(err, w)
		return
	}

	sendOK(appUser, w)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	appUser, err := user.CheckUser(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		sendError(err, w)
		return
	}

	tokenString, err := jwt.CreateToken(*appUser)
	if err != nil {
		sendError(err, w)
		return
	}

	rsp := map[string]string{
		"token": *tokenString,
	}

	sendOK(rsp, w)
}