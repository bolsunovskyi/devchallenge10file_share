package handlers

import (
	"github.com/gorilla/mux"
	"file_share/middleware"
	"net/http"
	"fmt"
	"net/http/httptest"
	"file_share/config"
	"file_share/test"
	"testing"
	"strings"
	"os"
)

func TestUploadFile(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", "foo.txt"),
		strings.NewReader("HELLO WORLD!"))
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	os.RemoveAll(config.Config.DataFolder)
}

func TestUploadFileWrongParent(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", "foo.txt"),
		strings.NewReader("HELLO WORLD!"))
	req.Header.Set("Access-Token", *token)
	req.Header.Set("File-Parent", "ffff")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}