package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"file_share/repository/file"
	"encoding/json"
	"file_share/models"
)

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

//TODO: add pagination
func ListFiles(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)
	var files []models.File
	var err error

	if parent, ok := vars["parent"]; ok {
		files, err = file.ListFiles(&parent, appUser)
	} else {
		files, err = file.ListFiles(nil, appUser)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(files)
}

func DeleteFile(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)
	err := file.DeleteFile(vars["fileID"], appUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func RenameFile(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)

	updateFile, err := file.RenameFile(vars["fileID"], r.FormValue("name"), appUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updateFile)
}