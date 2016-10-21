package user

import (
	"file_share/models"
	"gopkg.in/go-playground/validator.v9"
	"file_share/database"
	"gopkg.in/mgo.v2/bson"
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

func FindUserByEmail(email string) (*models.User, error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}

	user := models.User{}

	err = db.C(Collection).Find(bson.M{"email": email}).One(&user)
	session.Close()

	if err != nil {
		return nil, err
	}

	return &user, nil
}

//func DeleteUser()
