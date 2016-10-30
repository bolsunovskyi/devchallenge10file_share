package file

import (
	"testing"
	"file_share/config"
	"file_share/test"
	"file_share/repository/user"
	"os"
	"strings"
)

func TestUploadFile(t *testing.T) {
	defer test.TearDown(t)
	defer os.RemoveAll(config.Config.DataFolder)

	appUser, err := user.CreateUser("foo", "bar", "foo1@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = UploadFile(strings.NewReader("dqwdqwdw"), "dqwdqw.txt", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := CreateFolder("folder1", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	parent := folder1.ID.Hex()
	_, err = UploadFile(strings.NewReader("dqwdqwdw"), "dqwdqw.txt", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestUploadFileErrors(t *testing.T) {
	defer test.TearDown(t)
	defer os.RemoveAll(config.Config.DataFolder)

	port := config.Config.Mongo.Port
	config.Config.Mongo.Port = 64012

	if _, err := UploadFile(strings.NewReader("dqwdqwdw"), "dasdsad", nil, nil); err == nil {
		t.Error("No error on db err")
		return
	}

	config.Config.Mongo.Port = port

	appUser, err := user.CreateUser("foo", "bar", "foo1@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	parent := "sadasd"
	if _, err := UploadFile(strings.NewReader("dqwdqwdw"), "dasdsad", &parent, appUser); err == nil {
		t.Error("No error on wrong parent")
		return
	}

	if _, err := UploadFile(strings.NewReader("dqwdqwdw"), "name.txt", nil, appUser); err != nil {
		t.Error(err.Error())
		return
	}

	if _, err := UploadFile(strings.NewReader("dqwdqwdw"), "name.txt", nil, appUser); err == nil {
		t.Error("No error on same name")
		return
	}
}
