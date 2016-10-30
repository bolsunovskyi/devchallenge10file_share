package user

import (
	"file_share/models"
	"gopkg.in/go-playground/validator.v9"
	"file_share/database"
	"gopkg.in/mgo.v2/bson"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

//Collection var for user collection name
var Collection = "user"

//CreateUser creates user by credentials
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
	if err != nil {
		return nil, err
	}
	defer session.Close()

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

//FindUserByEmail looks for user by his email
func FindUserByEmail(email string) (*models.User, error) {
	session, db, err := database.GetSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	user := models.User{}

	err = db.C(Collection).Find(bson.M{"email": email}).One(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

//DeleteUser delete user by ID
func DeleteUser(userID bson.ObjectId) error {
	session, db, err := database.GetSession()
	if err != nil {
		return err
	}
	defer session.Clone()

	if err = db.C(Collection).RemoveId(userID); err != nil {
		return err
	}

	return nil
}

//CheckUser checks user email and password
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