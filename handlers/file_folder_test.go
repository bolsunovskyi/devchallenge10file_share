package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"file_share/middleware"
	"file_share/test"
	"net/http"
	"fmt"
	"net/http/httptest"
)

func TestCreateFolder(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
	}

	fileName := "images1"

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", fileName),
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()
	req.Header.Set("File-Folder", "true")
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		t.Error(recorder.Body)
	}
}

func TestCreateFolderWrongParent(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
	}

	fileName := "images1"

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", fileName),
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()
	req.Header.Set("File-Folder", "true")
	req.Header.Set("File-Parent", "foo")
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		t.Error(recorder.Body)
	}
}
