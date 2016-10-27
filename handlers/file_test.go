package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"os"
	"net/http"
	"net/http/httptest"
	"fmt"
	"file_share/test"
	"file_share/config"
	"file_share/middleware"
	"file_share/repository/user"
	"file_share/jwt"
)

func TestUploadFile(t *testing.T) {
	appUser, err := user.CreateUser("foo", "bar", "foo@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	token, err := jwt.CreateToken(*appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	fileName := "foo.txt"
	newFile, err := os.Create(fileName)
	if err != nil {
		t.Error(err.Error())
	}
	newFile.Write([]byte("Hello World"))
	newFile.Close()

	newFile, err = os.Open(fileName)
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", fileName),
		newFile)
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	newFile.Close()
	os.Remove(fileName)
	os.RemoveAll(config.Config.DataFolder)
	test.TearDown(t)
}

func TestCreateFolder(t *testing.T) {
	appUser, err := user.CreateUser("foo", "bar", "foo@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	token, err := jwt.CreateToken(*appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	fileName := "images"

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
	}

	test.TearDown(t)
}
