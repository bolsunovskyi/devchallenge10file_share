package handlers

import (
	"net/http"
	"file_share/repository/user"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := user.CreateUser(
		r.FormValue("first_name"),
		r.FormValue("last_name"),
		r.FormValue("email"),
		r.FormValue("password"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}