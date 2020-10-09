package database

import (
	"bank-app/helpers"
	"bank-app/users"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=bankapp sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func SetBalance(email string, newBalance int) {
	db := ConnectDB()
	var user users.User
	db.Where("email = ?", email).First(&user)
	user.Balence = newBalance

	db.Save(user)
}

func GetPayeeAndPayor(payeeEmail string, payorEmail string) (users.User, 
	.User) {
	db := ConnectDB()
	var payee users.User
	var payor users.User
	db.Where("email = ?", payeeEmail).First(&payee)
	db.Where("email = ?", payorEmail).First(&payor)

	return payee, payor
}

func TopUpAccountBalance(email string, amount int) {
	db := ConnectDB()
	var user users.User
	db.Where("email = ?", email).First(&user)
	startBalance := user.Balence
	topUpBalance := startBalance + amount
	user.Balence = topUpBalance

	db.Save(user)
}

func FindUser(email string) uint {
	db := ConnectDB()
	var user users.User
	//db.Table("users").Select("user_id").Where("email = ? ", email).First(&user.ID)
	db.Where("email = ?", email).First(&user)

	return user.ID
}

func GetUserBalance(email string) int {
	db := ConnectDB()
	var user users.User
	db.Where("email = ?", email).First(&user)

	return user.Balence
}

func IsUserPresent(email string) bool {
	userResult := FindUser(email)
	if userResult <= 0 {
		return false
	}
	return true
}

func CreateUser(username string, email string, password string) (users.User, bool) {
	//will move uuid later
	userID := uuid.Must(uuid.NewRandom())
	user := users.User{UserID: userID, Username: username, Email: email, Password: password}
	db := ConnectDB()
	db.AutoMigrate(&users.User{})
	db.Create(&user)
	// need to clean up returning true
	return user, true
}

func UpdateEmail(newEmail string, email string) {
	db := ConnectDB()
	var user users.User
	db.Where("email = ?", email).First(&user)
	user.Email = newEmail
	db.Save(user)

}
