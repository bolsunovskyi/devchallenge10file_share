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

func createUser() error {
	return CreateUser("Vasiliy", "Pupkin", "vasiliy@gmail.com", "123456");
}

func TestCreateUser(t *testing.T) {
	err := createUser()

	if err != nil {
		t.Error(err.Error())
	}
}

func TestFindUserByEmail(t *testing.T) {
	err := createUser()

	if err != nil {
		t.Error(err.Error())
	}

	_, err = FindUserByEmail("vasiliy@gmail.com")

	if err != nil {
		t.Error(err.Error())
	}
}
