package controllers

import (
	"cloud/app/service/auth/models"
	"cloud/cmd"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string) (string, error) {
	db, err := cmd.ConnectDB()
	if err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic(err)
	}

	userToken, err := GenToken(username)
	if err != nil {
		panic(err)
	}

	verifyErr := VerifyToken(userToken)
	if verifyErr != nil {
		panic(verifyErr)
	}

	db.Create(&models.User{
		Username: username,
		Password: hashedPassword,
	})

	return string(hashedPassword[:]), nil
}

func GenToken(username string) (string, error) {
	var secret = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	return signedToken, nil
}

func VerifyToken(token string) error {
	var secret = []byte(os.Getenv("SECRET_KEY"))
	verifiedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		panic(err)
	}

	if !verifiedToken.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ResetPassword() {
	return
}
