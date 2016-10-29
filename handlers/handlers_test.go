package handlers

import (
	"file_share/models"
	"file_share/repository/user"
	"file_share/jwt"
	"file_share/test"
)

func init() {
	test.InitConfig("../")
}

func createUserAndToken() (*models.User, *string, error) {
	appUser, err := user.CreateUser("foo", "bar", "foo8@gmail.com", "123456");
	if err != nil {
		return nil, nil, err
	}

	token, err := jwt.CreateToken(*appUser);
	if  err != nil {
		return nil, nil, err
	}

	return appUser, token, nil
}
