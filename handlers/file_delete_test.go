package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"fmt"
	"file_share/test"
	"file_share/middleware"
	"file_share/repository/user"
	"file_share/jwt"
	"file_share/repository/file"
	"gopkg.in/mgo.v2/bson"
)

func TestDeleteFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo7@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	token, err := jwt.CreateToken(*appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := file.CreateFolder("images23", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(DeleteFile))
	req, _ := http.NewRequest(
		"DELETE", fmt.Sprintf("/v1/file/%s", folder.ID.Hex()),
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestDeleteFileError(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(DeleteFile))
	req, _ := http.NewRequest(
		"DELETE", fmt.Sprintf("/v1/file/%s", bson.NewObjectId().Hex()),
		nil)
	req.Header.Set("Access-Token", *token)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
