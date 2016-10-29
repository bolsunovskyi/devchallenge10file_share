package middleware

import (
	"net/http"
	"encoding/json"
	"file_share/models"
	"file_share/jwt"
)

//Auth authentication middleware for handler functions
func Auth(f func(http.ResponseWriter, *http.Request, *models.User)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		if r.Header.Get("Access-Token") != "" {
			token = r.Header.Get("Access-Token")
		} else if r.FormValue("access-token") != "" {
			token = r.FormValue("access-token")
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Error{
				Message:        "Token is not passed",
			})
			return
		}

		if appUser, err := jwt.CheckToken(token); err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Error{
				Message:        err.Error(),
			})
			return
		} else {
			f(w, r, appUser)
		}
	})
}