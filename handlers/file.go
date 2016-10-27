package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"file_share/repository/file"
	"encoding/json"
	"file_share/models"
)

func createFolder(fileName string, parent *string, w http.ResponseWriter, appUser *models.User) {
	folder, err := file.CreateFolder(fileName, parent, appUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(folder)
}

func uploadFile(fileName string, parent *string, w http.ResponseWriter, r *http.Request, appUser *models.User) {
	uploadedFile, err := file.UploadFile(r.Body, fileName, parent, appUser);
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(uploadedFile)
}

func UploadFile(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	var parent *string = nil
	parentHeader := r.Header.Get("File-Parent")
	if parentHeader != "" {
		parent = &parentHeader
	}

	if folder := r.Header.Get("File-Folder"); folder != "" {
		createFolder(fileName, parent, w, appUser)
	} else {
		uploadFile(fileName, parent, w, r, appUser)
	}
}

func ListFiles(w http.ResponseWriter, r *http.Request, _ *models.User) {
	vars := mux.Vars(r)
	var files []models.File
	var err error

	if parent, ok := vars["parent"]; ok {
		files, err = file.ListFiles(&parent)
	} else {
		files, err = file.ListFiles(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}


	w.WriteHeader(200)
	json.NewEncoder(w).Encode(files)
}
