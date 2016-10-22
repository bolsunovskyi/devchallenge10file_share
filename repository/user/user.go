package user

import (
	"file_share/models"
	"gopkg.in/go-playground/validator.v9"
	"file_share/database"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var Collection string = "user"

func CreateUser(firstName string, lastName string, email string, password string) (*models.User, error) {
	user := models.User{
		FirstName:	firstName,
		LastName:	lastName,
		Email:		email,
		Password:	password,
	}

	if err := validator.New().Struct(user); err != nil {
		return nil, err
	}

	session, db, err := database.GetSession()
	defer session.Close()

	if err != nil {
		return nil, err
	}

	_, err = FindUserByEmail(email)

	if err == nil {
		return nil, errors.New("User already exists")
	}

	user.ID = bson.NewObjectId()
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(pass)
	db.C(Collection).Insert(&user)

	return &user, nil
}

func FindUserByEmail(email string) (*models.User, error) {
	session, db, err := database.GetSession()
	defer session.Close()

	if err != nil {
		return nil, err
	}

	user := models.User{}

	err = db.C(Collection).Find(bson.M{"email": email}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(userID bson.ObjectId) error {
	session, db, err := database.GetSession()
	defer session.Clone()

	if err != nil {
		return err
	}

	if err = db.C(Collection).RemoveId(userID); err != nil {
		return err
	}

	return nil
}

func CheckUser(email string, password string) (*models.User, error) {
	user, err := FindUserByEmail(email);
	if  err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}


	return user, nil
}