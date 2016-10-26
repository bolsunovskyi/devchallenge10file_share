package handlers

import (
	"net/http"
	"file_share/repository/user"
	"encoding/json"
	"file_share/jwt"
	"file_share/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := user.CreateUser(
		r.FormValue("first_name"),
		r.FormValue("last_name"),
		r.FormValue("email"),
		r.FormValue("password"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	appUser, err := user.CheckUser(r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	tokenString, err := jwt.CreateToken(*appUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	rsp := map[string]string{
		"token": *tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rsp)
}