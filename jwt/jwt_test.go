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
	if tokenUser, err := user.CreateUser("foo", "bar", "foo22@gmail.com", "123456"); err != nil {
		t.Error(err.Error())
	} else {
		if token, err := CreateToken(*tokenUser); err != nil {
			t.Error(err.Error())
		} else {
			if validUser, err := CheckToken(*token); err != nil {
				t.Error(err.Error())
			} else {
				if validUser.ID != tokenUser.ID {
					t.Error("User not equals")
					t.Error(validUser.ID)
					t.Error(tokenUser.ID)

					t.Error(validUser.Email)
					t.Error(tokenUser.Email)
				}
			}
		}
	}

	test.TearDown(t)
}
