package file

import (
	"testing"
	"file_share/test"
	"file_share/repository/user"
	"file_share/config"
)

func TestCreateFolder(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo2@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images1", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folderID := folder.ID.Hex()
	_, err = CreateFolder("summer2016", &folderID, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func TestCreateFolderErrors(t *testing.T) {
	defer test.TearDown(t)

	parent := "sadasd"
	if _, err := CreateFolder("asdasd", &parent, nil); err == nil {
		t.Error("No error on wrong parent")
		return
	}

	appUser, err := user.CreateUser("foo", "bar", "foo2@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = CreateFolder("images1", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = CreateFolder("images1", nil, appUser)
	if err == nil {
		t.Error("No error on same name")
		return
	}
}

func TestCreateFolderDBErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012

	if _, err := CreateFolder("dasdsad", nil, nil); err == nil {
		t.Error("No error on db err")
		return
	}

	config.Config.Mongo.Port = port
}
