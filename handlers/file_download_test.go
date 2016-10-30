package handlers

import (
	"net/http"
	"fmt"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
	"file_share/middleware"
	"gopkg.in/mgo.v2/bson"
	"file_share/test"
	"file_share/repository/file"
	"strings"
	"file_share/config"
)

func TestDownloadFileWrongID(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(DownloadFile))

	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/file/%s", "fooo"),
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

func TestDownloadFileUnExistID(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(DownloadFile))
	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/file/%s", bson.NewObjectId().Hex()),
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

func TestDownloadFileFolder(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}


	folder, err := file.CreateFolder("folderXXX", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	parent := folder.ID.Hex()
	_, err = file.UploadFile(strings.NewReader("sadasdas"), "asdsad", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	dataFolder := config.Config.DataFolder
	config.Config.DataFolder = "/dqdqwdqwdqwd"

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(DownloadFile))
	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/file/%s", folder.ID.Hex()),
		nil)
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	config.Config.DataFolder = dataFolder

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)

		return
	}

	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDownloadFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	appFile, err := file.UploadFile(strings.NewReader("Hello world!"), "fii.txt", nil, appUser)
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(DownloadFile))
	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/file/%s", appFile.ID.Hex()),
		nil)
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}
}
