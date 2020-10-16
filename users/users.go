package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateUser(username string, email string, password string) (interfaces.User, bool) {
	passwordHash := helpers.HashAndSalt([]byte(password))
	user, result := database.CreateUser(username, email, passwordHash)
	// need to clean up returning true
	return user, result
}

func PrepareToken(ID uint) string {
	tokenContent := jwt.MapClaims{
		"user_id": ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}
