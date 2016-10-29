package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"file_share/repository/file"
	"encoding/json"
	"file_share/models"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
	"io"
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
		sendError(err, w)
		return
	}

	sendOK(files, w)
}

func DeleteFile(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)

	err := file.DeleteFile(vars["fileID"], appUser)
	if err != nil {
		sendError(err, w)
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

	sendOK(updateFile, w)
}

func MoveFile(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)
	parentID := r.FormValue("parent_id")
	var parentPtr *string
	if parentID != "" {
		parentPtr = &parentID
	}

	updateFile, err := file.MoveFile(vars["fileID"], parentPtr, appUser)
	if err != nil {
		sendError(err, w)
		return
	}

	sendOK(updateFile, w)
}

//TODO: add pagination
func SearchFiles(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)
	files, err := file.SearchFiles(vars["keyword"], appUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error{
			Message: err.Error(),
		})
		return
	}

	sendOK(files, w)
}

func DownloadFile(w http.ResponseWriter, r *http.Request, appUser *models.User) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["fileID"]) {
		sendErrorStr("Wrong file ID", w)
		return
	}

	appFile, err := file.FindByIDUser(bson.ObjectIdHex(vars["fileID"]), appUser.ID)
	if err != nil {
		sendError(err, w)
		return
	}

	if appFile.IsDir {
		//TODO: add this ability
		sendErrorStr("You can not download folder", w)
		return
	}

	w.Header().Add("Content-Description", "File Transfer")
	w.Header().Add("Content-Transfer-Encoding", "binary")
	w.Header().Add("Connection", "Kepp-Alive")
	w.Header().Add("Pragma", "public")
	w.Header().Add("Content-Type", "application/force-download")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", appFile.Name))
	w.Header().Add("Content-Length", fmt.Sprintf("%d", appFile.FileSize))
	w.WriteHeader(http.StatusOK)


	filePath := fmt.Sprintf("%s/%s", appFile.RealPath, appFile.RealName)
	realFile, _ := os.Open(filePath)
	defer realFile.Close()

	io.Copy(w, realFile)
}