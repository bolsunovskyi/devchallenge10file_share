package handlers

import (
	"testing"
	"github.com/gorilla/mux"
	"os"
)

func TestUploadFile(t *testing.T) {
	fileName := "foo.txt"
	localFile, err := os.Create(fileName)
	if err != nil {
		t.Error(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/file/foo.txt", CreateUser)
}
