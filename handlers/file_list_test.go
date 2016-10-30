package handlers

import (
	"testing"
	"file_share/repository/file"
	"github.com/gorilla/mux"
	"file_share/middleware"
	"file_share/test"
	"net/http"
	"net/http/httptest"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func TestListFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = file.CreateFolder("images3", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/files", middleware.Auth(ListFiles))
	req, _ := http.NewRequest(
		"GET", "/v1/files",
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestListFilesParent(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := file.CreateFolder("images3", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/files/{parent}", middleware.Auth(ListFiles))
	req, _ := http.NewRequest(
		"GET", fmt.Sprintf("/v1/files/%s", folder1.ID.Hex()),
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestListFilesWrongParent(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/files/{parent}", middleware.Auth(ListFiles))
	req, _ := http.NewRequest(
		"GET", fmt.Sprintf("/v1/files/%s", bson.NewObjectId().Hex()),
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
