package user

import (
	"testing"
	"file_share/test"
)

func init() {
	test.InitConfig("../../")
}

func TestCreateUser(t *testing.T) {
	if _, err := CreateUser("Vasili1y", "Pupk1in", "vas1123iliy@gmail.com", "123456"); err != nil {
		t.Error(err)
	}

	test.TearDown(t)
}

func TestFindUserByEmail(t *testing.T) {
	_, err := CreateUser("Vasili1y", "Pupk1in", "vas1121123iliy@gmail.com", "123456");

	if  err != nil {
		t.Error(err.Error())
	} else {
		if _, err := FindUserByEmail("vas1121123iliy@gmail.com"); err != nil {
			t.Error(err.Error())
		}
	}

	test.TearDown(t)
}

func TestCheckUser(t *testing.T) {
	_, err := CreateUser("Vasili1y", "Pupk1in", "v11as1121123iliy@gmail.com", "123456");

	if  err != nil {
		t.Error(err.Error())
	} else {
		if _, err := CheckUser("v11as1121123iliy@gmail.com", "123456"); err != nil {
			t.Error(err.Error())
		}
	}

	test.TearDown(t)
}
