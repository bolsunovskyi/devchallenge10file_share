package jwt

import (
	"testing"
	"file_share/repository/user"
	"file_share/config"
	"fmt"
	"file_share/database"
)

func init() {
	config.File = "config_test.toml"
	if err := config.Read("../"); !err {
		fmt.Println("Unable to load config")
	}
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

	session, db, err := database.GetSession()
	defer  session.Close()

	if err != nil {
		t.Error(err.Error())
	}
	if err := db.DropDatabase(); err != nil {
		t.Error(err.Error())
	}
}
