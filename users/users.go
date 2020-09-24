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

func IsUserPresent(email string) bool {
	userResult := FindUser(email)
	if userResult <= 0 {
		return false
	}
	return true
}

func FindUser(email string) uint {
	db := helpers.ConnectDB()
	var user interfaces.User
	//db.Table("users").Select("user_id").Where("email = ? ", email).First(&user.ID)
	db.Where("email = ?", email).First(&user)

	return user.ID
}
