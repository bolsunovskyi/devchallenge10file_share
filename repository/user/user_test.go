package user

import (
	"testing"
	"fmt"
	"file_share/config"
	"file_share/models"
)

func init() {
	if err := config.Read("../../"); !err {
		fmt.Println("Unable to load config")
	}
}

func createUser() (*models.User, error) {
	return CreateUser("Vasiliy", "Pupkin", "vasiliy@gmail.com", "123456");
}

func TestCreateUser(t *testing.T) {
	if user, err := createUser(); err != nil {
		t.Error(err)
	} else {
		if err := DeleteUser(user.ID); err != nil {
			t.Error(err)
		}
	}
}

func TestFindUserByEmail(t *testing.T) {
	user, err := createUser();

	if  err != nil {
		t.Error(err.Error())
	}

	if _, err := FindUserByEmail("vasiliy@gmail.com"); err != nil {
		t.Error(err.Error())
	}

	if err = DeleteUser(user.ID); err != nil {
		t.Error(err.Error())
	}
}

func TestCheckUser(t *testing.T) {
	user, err := createUser();

	if  err != nil {
		t.Error(err.Error())
	}

	if err := CheckUser("vasiliy@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	}

	if err = DeleteUser(user.ID); err != nil {
		t.Error(err.Error())
	}
}
