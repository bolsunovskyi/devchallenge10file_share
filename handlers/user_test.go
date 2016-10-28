package handlers

import (
	"testing"
	"net/http"
	"strings"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"encoding/json"
	"file_share/models"
	"file_share/repository/user"
	"file_share/test"
)

func init() {
	test.InitConfig("../")
}

func TestCreateUser(t *testing.T) {
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
	}

	appUser := models.User{}
	if err := json.NewDecoder(recorder.Body).Decode(&appUser); err != nil {
		t.Error(err.Error())
	}

	test.TearDown(t)
}

func TestLoginUser(t *testing.T) {
	if _, err := user.CreateUser("foo", "bar", "foo12@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
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
		}

		tokenMap := make(map[string]string)
		if err := json.NewDecoder(recorder.Body).Decode(&tokenMap); err != nil {
			t.Error("Unable to parse respnse")
		}
	}

	test.TearDown(t)
}