package handlers

import (
	"net/http"
)

//Front dummy handler
func Front(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
