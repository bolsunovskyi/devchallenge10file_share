package file

import (
	"testing"
	"file_share/config"
	"file_share/test"
	"file_share/repository/user"
	"gopkg.in/mgo.v2/bson"
)

func TestFindByIDDBErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012
	_, err := FindByID(bson.NewObjectId())
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	config.Config.Mongo.Port = port
}

func TestFindByNameDBErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012
	_, err := FindByNameAndDir("ssss", nil)
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	config.Config.Mongo.Port = port
}

func TestFindByName(t *testing.T) {
	defer test.TearDown(t)

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

	_, err = FindByNameAndDir("images1", nil)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestFindByIDUserDBErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012
	_, err := FindByIDUser(bson.NewObjectId(), bson.NewObjectId())
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	config.Config.Mongo.Port = port
}

func TestFindByIDWrongID(t *testing.T) {
	if _, err := FindByID(bson.NewObjectId()); err == nil {
		t.Error("No error on wrong id")
	}
}

func TestFindByIDUserWrongID(t *testing.T) {
	if _, err := FindByIDUser(bson.NewObjectId(), bson.NewObjectId()); err == nil {
		t.Error("No error on wrong id")
	}
}
