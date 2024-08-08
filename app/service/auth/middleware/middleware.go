package middleware

import (
	"cloud/app/service/auth/models"
	"cloud/cmd"

	"golang.org/x/crypto/bcrypt"
)

func VerifyUser(username string, password string) bool {
	err := cmd.LoadEnv()
	if err != nil {
		panic(err)
	}
	db, dberr := cmd.ConnectDB()
	if dberr != nil {
		panic(err)
	}
	var user models.User
	db.First(&user, "username = ?", username)
	compareErr := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if compareErr != nil {
		return false
	}
	return true
}
