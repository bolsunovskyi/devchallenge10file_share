package handlers

import (
	"net/http"
	"encoding/json"
	"file_share/models"
)

func sendErrorStr(msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.Error{
		Message: msg,
	})
}

func sendError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.Error{
		Message: err.Error(),
	})
}

func sendOK(v interface {}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
