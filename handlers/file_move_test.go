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
	"gopkg.in/mgo.v2/bson"
)

func TestMoveFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := file.CreateFolder("images3123", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	folder2, err := file.CreateFolder("imag234es3123", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(MoveFile))

	req, _ := http.NewRequest(
		"PATCH",  fmt.Sprintf("/v1/file/%s", folder1.ID.Hex()),
		strings.NewReader(fmt.Sprintf("parent_id=%s", folder2.ID.Hex())))
	req.Header.Set("Access-Token", *token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestMoveFileWrongID(t *testing.T) {
	defer test.TearDown(t)

	_, token, err := createUserAndToken()
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileID}", middleware.Auth(MoveFile))

	req, _ := http.NewRequest(
		"PATCH",  fmt.Sprintf("/v1/file/%s", bson.NewObjectId().Hex()),
		strings.NewReader(fmt.Sprintf("parent_id=%s", bson.NewObjectId().Hex())))
	req.Header.Set("Access-Token", *token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
