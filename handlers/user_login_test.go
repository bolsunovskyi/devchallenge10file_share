package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"net/http/httptest"
	"encoding/json"
	"file_share/test"
	"file_share/repository/user"
)

func TestLoginUser(t *testing.T) {
	defer test.TearDown(t)

	_, err := user.CreateUser("foo", "bar", "foo12@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/auth", LoginUser)
	req, _ := http.NewRequest(
		"POST", "/v1/auth",
		strings.NewReader("email=foo12@gmail.com&password=123456"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		return
	}

	tokenMap := make(map[string]string)
	if err := json.NewDecoder(recorder.Body).Decode(&tokenMap); err != nil {
		t.Error("Unable to parse respnse")
	}
}

func TestLoginUserWrong(t *testing.T) {
	defer test.TearDown(t)

	router := mux.NewRouter()
	router.HandleFunc("/v1/auth", LoginUser)
	req, _ := http.NewRequest(
		"POST", "/v1/auth",
		strings.NewReader("email=foo12@gmail.com&password=123456"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
		return
	}
}
