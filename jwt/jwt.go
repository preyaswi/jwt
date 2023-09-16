package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var Mysigninkey = []byte("qwertyuiop")

func GenerateToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorised"] = true
	claims["user"] = "Preya"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(Mysigninkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
