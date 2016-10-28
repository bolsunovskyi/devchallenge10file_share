package handlers

import (
	"net/http"
	"file_share/repository/file"
	"file_share/models"
	"encoding/json"
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
