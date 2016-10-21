package models

type User struct {
	FirstName	string		`validate:"required"`
	LastName	string		`validate:"required"`
	Email		string		`validate:"required"`
	Password	string		`validate:"required"`
	Token		[]string
}
