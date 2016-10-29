package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
)

func TestFront(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/", Front)

	req, _ := http.NewRequest(
		"GET",  "/",
		nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}
