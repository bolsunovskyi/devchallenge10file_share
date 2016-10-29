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
	"file_share/middleware"
	"file_share/repository/user"
	"file_share/jwt"
	"file_share/repository/file"
	"strings"
	"encoding/json"
	"file_share/models"
)

func TestUploadFile(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo5@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if token, err := jwt.CreateToken(*appUser); err != nil {
			t.Error(err.Error())
		} else {
			fileName := "foo.txt"
			if newFile, err := os.Create(fileName); err != nil {
				t.Error(err.Error())
			} else {
				newFile.Write([]byte("Hello World"))
				newFile.Close()

				newFile, _ = os.Open(fileName)

				router := mux.NewRouter()
				router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

				req, _ := http.NewRequest(
					"POST", fmt.Sprintf("/v1/file/%s", fileName),
					newFile)
				req.Header.Set("Access-Token", *token)

				recorder := httptest.NewRecorder()
				router.ServeHTTP(recorder, req)

				if status := recorder.Code; status != http.StatusOK {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, http.StatusOK)
				}

				newFile.Close()
				os.Remove(fileName)
				os.RemoveAll(config.Config.DataFolder)
			}
		}
	}

	test.TearDown(t)
}

func TestCreateFolder(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo6@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if token, err := jwt.CreateToken(*appUser); err != nil {
			t.Error(err.Error())
		} else {
			fileName := "images1"

			router := mux.NewRouter()
			router.HandleFunc("/v1/file/{fileName:[0-9a-zA-Z._]+}", middleware.Auth(UploadFile))

			req, _ := http.NewRequest(
				"POST", fmt.Sprintf("/v1/file/%s", fileName),
				nil)
			req.Header.Set("Access-Token", *token)
			recorder := httptest.NewRecorder()
			req.Header.Set("File-Folder", "true")
			router.ServeHTTP(recorder, req)

			if status := recorder.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
				t.Error(recorder.Body)
			}
		}
	}

	test.TearDown(t)
}

func TestDeleteFile(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo7@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if token, err := jwt.CreateToken(*appUser); err != nil {
			t.Error(err.Error())
			return
		} else {
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
	}

	test.TearDown(t)
}

func TestRenameFile(t *testing.T) {
	if appUser, err := user.CreateUser("foo", "bar", "foo8@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if token, err := jwt.CreateToken(*appUser); err != nil {
			t.Error(err.Error())
		} else {
			if folder, err := file.CreateFolder("images3", nil, appUser); err != nil {
				t.Error(err.Error())
			} else {
				router := mux.NewRouter()
				router.HandleFunc("/v1/file/{fileID}", middleware.Auth(RenameFile))
				req, _ := http.NewRequest(
					"PUT", fmt.Sprintf("/v1/file/%s", folder.ID.Hex()),
					strings.NewReader("name=bar"))
				req.Header.Set("Access-Token", *token)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				recorder := httptest.NewRecorder()

				router.ServeHTTP(recorder, req)

				if status := recorder.Code; status != http.StatusOK {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, http.StatusOK)
				}
			}
		}
	}

	test.TearDown(t)
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

func TestSearchFiles(t *testing.T) {
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

	_, err = file.CreateFolder("images3123", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	router := mux.NewRouter()
	router.HandleFunc("/v1/files/search/{keyword}", middleware.Auth(SearchFiles))

	req, _ := http.NewRequest(
		"GET",  fmt.Sprintf("/v1/files/search/%s", "image"),
		nil)
	req.Header.Set("Access-Token", *token)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var files []models.File
	json.NewDecoder(recorder.Body).Decode(&files)

	if len(files) == 0 {
		t.Error("No files found")
	}
}