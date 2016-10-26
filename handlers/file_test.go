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
)

func TestUploadFile(t *testing.T) {
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
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", UploadFile)

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", fileName),
		newFile)
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
	fileName := "images"

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", UploadFile)

	req, _ := http.NewRequest(
		"POST", fmt.Sprintf("/v1/file/%s", fileName),
		nil)
	recorder := httptest.NewRecorder()
	req.Header.Set("File-Folder", "true")
	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	test.TearDown(t)
}
