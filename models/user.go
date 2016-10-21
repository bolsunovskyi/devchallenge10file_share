package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID		bson.ObjectId	`bson:"_id,omitempty"`
	FirstName	string		`validate:"required" bson:"first_name"`
	LastName	string		`validate:"required"`
	Email		string		`validate:"required,email"`
	Password	string		`validate:"required"`
}

type LoginUser struct {
	ID		bson.ObjectId	`bson:"_id,omitempty"`
	Email		string		`validate:"required,email"`
	Password	string		`validate:"required"`
	Token		string
}
