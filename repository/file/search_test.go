package file

import (
	"testing"
	"file_share/test"
	"file_share/repository/user"
	"file_share/config"
)

func TestSearchFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo2221@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = CreateFolder("images", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err := SearchFiles("imag", appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) == 0 {
		t.Error("No files found")
		return
	}
}

func TestSearchFilesError(t *testing.T) {
	port := config.Config.Mongo.Port
	config.Config.Mongo.Port = 64012

	if _, err := SearchFiles("", nil); err == nil {
		t.Error("No error on db err")
		return
	}

	config.Config.Mongo.Port = port

	if _, err := SearchFiles("", nil); err == nil {
		t.Error("No error on short keyword")
		return
	}
}