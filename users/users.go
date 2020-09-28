package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
)

func CreateUser(username string, email string, password string) (interfaces.User, bool) {
	passwordHash := helpers.HashAndSalt([]byte(password))
	user, result := database.CreateUser(username, email, passwordHash)
	// need to clean up returning true
	return user, result
}
