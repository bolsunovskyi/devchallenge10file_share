package handlers

import (
	"testing"
	"net/http"
	"strings"
	"net/http/httptest"
	"github.com/gorilla/mux"
	"fmt"
	"file_share/config"
	"encoding/json"
	"file_share/models"
	"file_share/repository/user"
	"file_share/database"
)

func init() {
	config.File = "config_test.toml"
	if err := config.Read("../"); !err {
		fmt.Println("Unable to load config")
	}
}

func down(t *testing.T) {
	session, db, err := database.GetSession()
	defer  session.Close()

	if err != nil {
		t.Error(err.Error())
	}
	if err := db.DropDatabase(); err != nil {
		t.Error(err.Error())
	}
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

	if err := user.DeleteUser(appUser.ID); err != nil {
		t.Error(err.Error())
	}

	down(t)
}

func TestLoginUser(t *testing.T) {
	_, err := user.CreateUser("foo", "bar", "foo@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/auth", LoginUser)
	req, _ := http.NewRequest(
		"POST", "/v1/auth",
		strings.NewReader("email=foo@gmail.com&password=123456"))
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
	down(t)
}