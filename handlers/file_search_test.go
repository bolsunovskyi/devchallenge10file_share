package handlers

import (
	"file_share/repository/file"
	"github.com/gorilla/mux"
	"file_share/middleware"
	"file_share/test"
	"net/http"
	"fmt"
	"net/http/httptest"
	"file_share/models"
	"encoding/json"
	"testing"
)

func TestSearchFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()

	_, err = file.CreateFolder("images3123", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/files/search/{keyword}", middleware.Auth(SearchFiles))

	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/files/search/%s", "image"),
		nil)
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	var files []models.File
	json.NewDecoder(recorder.Body).Decode(&files)

	if len(files) == 0 {
		t.Error("No files found")
	}
}

func TestSearchFilesEmptyKeyword(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()

	_, err = file.CreateFolder("images3123", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/files/search/{keyword}", middleware.Auth(SearchFiles))

	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/files/search/%s", "a"),
		nil)
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		return
	}
}
