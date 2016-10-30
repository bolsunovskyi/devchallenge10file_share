package file

import (
	"testing"
	"file_share/test"
	"file_share/repository/user"
	"gopkg.in/mgo.v2/bson"
	"file_share/config"
)

func TestMoveFile(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo1255@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := CreateFolder("images223s4", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	folder2, err := CreateFolder("image23rs223s4", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	parentID := folder2.ID.Hex()
	_, err = MoveFile(folder1.ID.Hex(), &parentID, appUser)
	if  err != nil {
		t.Error(err.Error())
		return
	}

	_, err = MoveFile(folder1.ID.Hex(), nil, appUser)
	if  err != nil {
		t.Error(err.Error())
		return
	}
}

func TestMoveFileWrongID(t *testing.T) {
	defer test.TearDown(t)

	if _, err := MoveFile("asdsad", nil, nil); err == nil {
		t.Error("No error on wrong id")
		return
	}

	appUser, err := user.CreateUser("foo", "bar", "foo1255@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	parent := bson.NewObjectId().Hex()
	if _, err := MoveFile(bson.NewObjectId().Hex(), &parent, appUser); err == nil {
		t.Error("No error on wrong id")
		return
	}
}

func TestMoveFileDBErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012

	parent := bson.NewObjectId().Hex()
	if _, err := MoveFile(bson.NewObjectId().Hex(), &parent, nil); err == nil {
		t.Error("No error on db err")
		return
	}

	config.Config.Mongo.Port = port
}

func TestMoveFileWrongFileID(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := user.CreateUser("foo", "bar", "foo1255@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	folder1, err := CreateFolder("images223s4", nil, appUser);
	if  err != nil {
		t.Error(err.Error())
		return
	}

	parent := folder1.ID.Hex()
	if _, err := MoveFile(bson.NewObjectId().Hex(), &parent, appUser); err == nil {
		t.Error("No error on db err")
		return
	}
}


