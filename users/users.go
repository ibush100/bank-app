package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, email string, password string) (users.User, bool) {
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

func checkPass(email string, password string) bool {
	db := database.ConnectDB()
	var user users.User
	db.Where("email = ?", email).First(&user)
	passCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if passCheck == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}
