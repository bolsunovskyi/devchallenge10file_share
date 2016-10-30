package models

import "gopkg.in/mgo.v2/bson"

//User type for representing user entity
type User struct {
	ID		bson.ObjectId	`bson:"_id,omitempty" json:"id"`
	FirstName	string		`validate:"required" bson:"first_name" json:"first_name"`
	LastName	string		`validate:"required" json:"last_name"`
	Email		string		`validate:"required,email" json:"email"`
	Password	string		`validate:"required,min=3" json:"-"`
}

//LoginUser type for validating user login
type LoginUser struct {
	ID		bson.ObjectId	`bson:"_id,omitempty"`
	Email		string		`validate:"required,email"`
	Password	string		`validate:"required"`
	Token		string
}
