package user

import (
	"testing"
	"file_share/test"
)

func TestCreateUser(t *testing.T) {
	defer test.TearDown(t)

	if _, err := CreateUser("Vasili1y", "Pupk1in", "vas1123iliy@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	}
}

func TestCreateUserWrongEmail(t *testing.T) {
	if _, err := CreateUser("Vasili1y", "Pupk1in", "vas1123iliy", "123456"); err == nil {
		t.Error("No error on wrong email")
	}
}

func  TestCreateUserExists(t *testing.T) {
	defer test.TearDown(t)

	if _, err := CreateUser("Vasili1y", "Pupk1in", "vas1123iliy@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	}

	if _, err := CreateUser("Vasili1y", "Pupk1in", "vas1123iliy@gmail.com", "123456"); err == nil {
		t.Error("No error on user exists")
	}
}
