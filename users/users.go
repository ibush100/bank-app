package users

import (
	"bank-app/helpers"
	"bank-app/interfaces"

	"github.com/google/uuid"
)

func CreateUser(username string, email string, password string) (interfaces.User, bool) {
	userID := uuid.Must(uuid.NewRandom())
	passwordHash := helpers.HashAndSalt([]byte(password))
	user := interfaces.User{UserID: userID, Username: username, Email: email, Password: passwordHash}
	db := helpers.ConnectDB()
	db.AutoMigrate(&interfaces.User{})
	db.Create(&user)
	// need to clean up returning true
	return user, true
}
