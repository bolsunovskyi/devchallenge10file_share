package user

import (
	"testing"
	"file_share/models"
	"file_share/test"
)

func init() {
	test.InitConfig("../../")
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

	test.TearDown(t)
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

	test.TearDown(t)
}

func TestCheckUser(t *testing.T) {
	user, err := createUser();

	if  err != nil {
		t.Error(err.Error())
	}

	if _, err := CheckUser("vasiliy@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	}

	if err = DeleteUser(user.ID); err != nil {
		t.Error(err.Error())
	}

	test.TearDown(t)
}
