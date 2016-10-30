package file

import (
	"testing"
	"file_share/repository/user"
	"file_share/test"
	"gopkg.in/mgo.v2/bson"
	"file_share/config"
)

func TestRenameFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo5@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images4", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = RenameFile(folder.ID.Hex(), "photos", appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	appFile, err := FindByID(folder.ID)
	if  err != nil {
		t.Error(err.Error())
		return
	}
	if appFile.Name != "photos" {
		t.Error("File is not updated")
	}
}

func TestRenameFileErrors(t *testing.T) {
	defer test.TearDown(t)

	if _, err := RenameFile("ccsacas", "sadasd", nil); err == nil {
		t.Error("No error on wrong file id")
		return
	}

	if _, err := RenameFile(bson.NewObjectId().Hex(), "!@?sadasd", nil); err == nil {
		t.Error("No error on wrong file name")
		return
	}

	port := config.Config.Mongo.Port
	config.Config.Mongo.Port = 64012

	if _, err := RenameFile(bson.NewObjectId().Hex(), "normal_name", nil); err == nil {
		t.Error("No error on db err")
		return
	}

	config.Config.Mongo.Port = port


	appUser, err := user.CreateUser("foo", "bar", "foo5@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if _, err := RenameFile(bson.NewObjectId().Hex(), "DQWDQdasd", appUser); err == nil {
		t.Error("No error on UNEXIST file id")
		return
	}
}