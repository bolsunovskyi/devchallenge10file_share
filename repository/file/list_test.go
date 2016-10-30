package file

import (
	"testing"
	"file_share/test"
	"file_share/repository/user"
	"file_share/config"
)

func TestListFiles(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo3@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := CreateFolder("images2", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err := ListFiles(nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) == 0 {
		t.Error("Empty file list")
	}

	parent := folder1.ID.Hex()
	_, err = CreateFolder("images23", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err = ListFiles(&parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) == 0 {
		t.Error("Empty file list")
	}
}

func TestListFilesError(t *testing.T) {
	port := config.Config.Mongo.Port
	config.Config.Mongo.Port = 64012

	if _, err := ListFiles(nil, nil); err == nil {
		t.Error("No error on db err")
		return
	}

	config.Config.Mongo.Port = port
}
