package user

import (
	"file_share/models"
	"gopkg.in/go-playground/validator.v9"
	"file_share/database"
)

var Collection string = "user"

func CreateUser(firstName string, lastName string, email string, password string) error {
	user := models.User{
		FirstName:	firstName,
		LastName:	lastName,
		Email:		email,
		Password:	password,
	}

	if err := validator.New().Struct(user); err != nil {
		return err
	}

	session, db, err := database.GetSession()
	if err != nil {
		return err
	}

	//TODO: add check for user exists
	db.C(Collection).Insert(&user)
	session.Close()

	return nil
}
