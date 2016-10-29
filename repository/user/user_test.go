package user

import (
	"testing"
	"file_share/test"
)

func init() {
	test.InitConfig("../../")
}

func TestCreateUser(t *testing.T) {
	defer test.TearDown(t)

	if _, err := CreateUser("Vasili1y", "Pupk1in", "vas1123iliy@gmail.com", "123456"); err != nil {
		t.Error(err)
	}
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
	}
}