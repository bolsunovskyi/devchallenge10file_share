package jwt

import (
	"file_share/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"file_share/config"
	"fmt"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

//CreateToken creates new access token for specified user
func CreateToken(tokenUser models.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID": tokenUser.ID.Hex(),
		"FirstName": tokenUser.FirstName,
		"LastName": tokenUser.LastName,
		"Email": tokenUser.Email,
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(config.Config.Secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

//CheckToken verifies token and returns user
func CheckToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		returnUser := models.User{
			ID: 		bson.ObjectIdHex(claims["ID"].(string)),
			Email:		claims["Email"].(string),
			Password:	"",
			FirstName:	claims["FirstName"].(string),
			LastName:	claims["LastName"].(string),
		}

		return &returnUser, nil
	}

	return nil, errors.New("Unable to validate token")
}
