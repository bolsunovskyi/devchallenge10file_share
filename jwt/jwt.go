package jwt

import (
	"file_share/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"file_share/config"
)

func CreateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"nbf": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(config.Config.Secret)
	if err != nil {
		return nil, err
	}

	return tokenString, nil
}

func CheckToken(token string) (*models.User, error) {
	return nil, nil
}
