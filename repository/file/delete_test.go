package file

import (
	"testing"
	"file_share/test"
	"file_share/repository/user"
	"file_share/config"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"os"
)

func TestDeleteFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images3", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	folderID := folder.ID.Hex()
	err = DeleteFile(folderID, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	files, err := ListFiles(nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if len(files) > 0 {
		t.Error("File list is not empty")
	}
}

func TestDeleteFileWrongIDFormat(t *testing.T) {
	err := DeleteFile("sss", nil)
	if err == nil {
		t.Error("No error on wrong id")
		return
	}
}

func TestDeleteFileConnErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012

	err := DeleteFile(bson.NewObjectId().Hex(), nil)
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	config.Config.Mongo.Port = port
}

func TestDeleteFileWrongID(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = DeleteFile(bson.NewObjectId().Hex(), appUser)
	if err == nil {
		t.Error("No error on wrong file id")
		return
	}
}

func TestDeleteFileWrongUser(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "fowqdqwdo4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images3", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	appUser2, err := user.CreateUser("foo", "bar", "foasdsado4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = DeleteFile(folder.ID.Hex(), appUser2)
	if err == nil {
		t.Error("No error on wrong user")
		return
	}
}

func TestDeleteRealFile(t *testing.T) {
	defer test.TearDown(t)
	defer os.RemoveAll(config.Config.DataFolder)

	appUser, err := user.CreateUser("foo", "bar", "fowqdqwdo4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	uFile, err := UploadFile(strings.NewReader("fooo baar"), "sadasd", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = DeleteFile(uFile.ID.Hex(), appUser)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestRecursiveDelete(t *testing.T) {
	defer test.TearDown(t)
	defer os.RemoveAll(config.Config.DataFolder)

	appUser, err := user.CreateUser("foo", "bar", "fowqdqwdo4@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	folder, err := CreateFolder("images3", nil, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	parent := folder.ID.Hex()
	folder1, err := CreateFolder("images3", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	_, err = UploadFile(strings.NewReader("fooo baar"), "sadasd", &parent, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	parent2 := folder1.ID.Hex()
	_, err = UploadFile(strings.NewReader("fooo baar"), "sadasd", &parent2, appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = DeleteFile(folder.ID.Hex(), appUser)
	if err != nil {
		t.Error(err.Error())
		return
	}
}
