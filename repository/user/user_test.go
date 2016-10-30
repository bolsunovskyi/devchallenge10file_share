package user

import (
	"testing"
	"file_share/test"
	"file_share/config"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	test.InitConfig("../../")
}

func TestFindUserByEmail(t *testing.T) {
	defer test.TearDown(t)

	_, err := CreateUser("Vasili1y", "Pupk1in", "vas1121123iliy@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	if _, err := FindUserByEmail("vas1121123iliy@gmail.com"); err != nil {
		t.Error(err.Error())
	}
}

func TestCheckUser(t *testing.T) {
	defer test.TearDown(t)

	_, err := CreateUser("Vasili1y", "Pupk1in", "v11as1121123iliy@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	_, err = CheckUser("v11as1121123iliy@gmail.com", "123456")
	if err != nil {
		t.Error(err.Error())
		return
	}

	if _, err := CheckUser("dqwdqw", "qwdqw"); err == nil {
		t.Error("No error on wrong email and password")
		return
	}

	_, err = CheckUser("v11as1121123iliy@gmail.com", "1234567")
	if err == nil {
		t.Error("No error on wrong password")
		return
	}
}

func TestDeleteUser(t *testing.T) {
	defer test.TearDown(t)

	appUser, err := CreateUser("Vasili1y", "Pupk1in", "v11as112ssd1123iliy@gmail.com", "123456");
	if  err != nil {
		t.Error(err.Error())
		return
	}

	err = DeleteUser(appUser.ID)
	if err != nil {
		t.Error(err.Error())
		return
	}

	err = DeleteUser(bson.NewObjectId())
	if err == nil {
		t.Error("No error on wrong id")
		return
	}
}

func TestFindUserDBErr(t *testing.T) {
	port := config.Config.Mongo.Port

	config.Config.Mongo.Port = 64012
	_, err := FindUserByEmail("sadasd")
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	_, err = CreateUser("sadas", "sdasd", "dqwdwq@sadasd.cc", "dassd")
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	err = DeleteUser(bson.NewObjectId())
	if err == nil {
		t.Error("No error on wrong mongo port")
	}

	config.Config.Mongo.Port = port
}