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
	"strings"
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



func TestListFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo8@gmail.com", "123456");
	if err != nil {
		t.Error(err.Error())
		return
	}

	token, err := jwt.CreateToken(*appUser);
	if  err != nil {
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

func TestMoveFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo8@gmail.com", "123456");
	if err != nil {
		t.Error(err.Error())
		return
	}

	token, err := jwt.CreateToken(*appUser);
	if  err != nil {
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
	router.HandleFunc("/v1/file/move/{fileID}", middleware.Auth(MoveFile))

	req, _ := http.NewRequest(
		"POST",  fmt.Sprintf("/v1/file/move/%s", folder1.ID.Hex()),
		strings.NewReader(fmt.Sprintf("parentID=%s", folder2.ID.Hex())))
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}



