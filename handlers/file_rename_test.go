package handlers

import (
	"testing"
	"file_share/repository/file"
	"github.com/gorilla/mux"
	"file_share/middleware"
	"file_share/test"
	"net/http"
	"fmt"
	"strings"
	"net/http/httptest"
)

func TestRenameFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
	}

	folder, err := file.CreateFolder("images3", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(RenameFile))
	req, _ := http.NewRequest(
		"PUT", fmt.Sprintf("/v1/file/%s", folder.ID.Hex()),
		strings.NewReader("name=bar"))
	req.Header.Set("Access-Token", *token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestRenameFileWrongID(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(RenameFile))
	req, _ := http.NewRequest(
		"PUT", fmt.Sprintf("/v1/file/%s", "fffff"),
		strings.NewReader("name=bar"))
	req.Header.Set("Access-Token", *token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

