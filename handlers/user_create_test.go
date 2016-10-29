package handlers

import (
	"net/http/httptest"
	"net/http"
	"file_share/models"
	"file_share/test"
	"encoding/json"
	"testing"
	"github.com/gorilla/mux"
	"strings"
	"file_share/repository/user"
)

func TestCreateUser(t *testing.T) {
	defer test.TearDown(t)

	router := mux.NewRouter()
	router.HandleFunc("/v1/user", CreateUser)
	req, _ := http.NewRequest(
		"POST", "/v1/user",
		strings.NewReader("first_name=vasia&last_name=pupkin&email=vasia@gmail.com&password=123456"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	appUser := models.User{}
	if err := json.NewDecoder(recorder.Body).Decode(&appUser); err != nil {
		t.Error(err.Error())
	}
}

func TestCreateUserErr(t *testing.T) {
	defer test.TearDown(t)

	_, err := user.CreateUser("foo", "asd", "exit2@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/user", CreateUser)
	req, _ := http.NewRequest(
		"POST", "/v1/user",
		strings.NewReader("first_name=vasia&last_name=pupkin&email=exit2@gmail.com&password=123456"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		return
	}
}