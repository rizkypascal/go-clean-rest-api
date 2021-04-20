package services

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func getHash(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}
