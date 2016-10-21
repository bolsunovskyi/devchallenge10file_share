package user

import (
	"testing"
	"fmt"
	"file_share/config"
)

func init() {
	if err := config.Read("../../"); !err {
		fmt.Println("Unable to load config")
	}
}

func TestCreateUser(t *testing.T) {
	err := CreateUser("Vasiliy", "Pupkin", "vasiliy@gmail.com", "123456");

	if err != nil {
		t.Error(err.Error())
		t.Error(fmt.Sprintf("%s:%d", config.Config.Mongo.Host, config.Config.Mongo.Port))
	}
}
