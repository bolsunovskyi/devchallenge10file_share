package jwt

import (
	"testing"
	"file_share/repository/user"
	"file_share/test"
)

func init() {
	test.InitConfig("../")
}

func TestCreateToken(t *testing.T) {
	tokenUser, err := user.CreateUser("foo", "bar", "foo@gmail.com", "123456")
	if(err != nil) {
		t.Error(err.Error())
	}

	token, err := CreateToken(*tokenUser)
	if err != nil {
		t.Error(err.Error())
	} else {

		validUser, err := CheckToken(*token)
		if err != nil {
			t.Error(err.Error())
		}

		if validUser.ID != tokenUser.ID {
			t.Error("User not equals")
			t.Error(validUser.ID)
			t.Error(tokenUser.ID)

			t.Error(validUser.Email)
			t.Error(tokenUser.Email)
		}
	}

	if err := user.DeleteUser(tokenUser.ID); err != nil {
		t.Error(err.Error())
	}

	test.TearDown(t)
}
