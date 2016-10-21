package handlers

import (
	"net/http"
	"file_share/repository/user"
	"encoding/json"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.CreateUser(
		r.FormValue("first_name"),
		r.FormValue("last_name"),
		r.FormValue("email"),
		r.FormValue("password"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}