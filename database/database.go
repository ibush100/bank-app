package database

import (
	"bank-app/helpers"
	"bank-app/interfaces"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 dbname=bankapp sslmode=disable")
	helpers.HandleErr(err)
	return db
}

func SetBalance(email string, newBalance int) {
	db := ConnectDB()
	account := GetAccount(email)
	account.Balance = newBalance
	db.Save(account)
}

func CreateAccount(email string) interfaces.Account {
	userID := FindUser(email)
	account := interfaces.Account{OwnerID: userID}
	db := ConnectDB()
	db.Create(&account)
	return account
}

func GetAccount(email string) interfaces.Account {
	userID := FindUser(email)
	account := interfaces.Account{OwnerID: userID}
	db := ConnectDB()
	db.Where("OwnerID = ?", userID).First(&account)
	return account
}

func GetPayeeAndPayor(payeeEmail string, payorEmail string) (interfaces.User, interfaces.User) {
	db := ConnectDB()
	var payee interfaces.User
	var payor interfaces.User
	db.Where("email = ?", payeeEmail).First(&payee)
	db.Where("email = ?", payorEmail).First(&payor)

	return payee, payor
}

func TopUpAccountBalance(email string, amount int) {
	db := ConnectDB()
	account := GetAccount(email)
	startBalance := account.Balance
	topUpBalance := startBalance + amount
	account.Balance = topUpBalance

	db.Save(account)
}

func FindUser(email string) uuid.UUID {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)

	return user.UserID
}

func FindUserEmail(email string) string {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	return user.Email
}

func IsUserEmailAvailable(email string) bool {
	db := ConnectDB()
	var user interfaces.User
	if result := db.Where("email = ?", email).First(&user); result.Error != nil {
		return false
	}
	return true
}

func GetUserBalance(email string) int {
	account := GetAccount(email)
	return account.Balance
}

func IsUserPresent(email string) bool {
	userResult := FindUser(email)
	if userResult == uuid.Nil {
		return false
	}
	return true
}

func CreateUser(username string, email string, password string) (interfaces.User, bool) {
	var result bool
	emailPresent := IsUserEmailAvailable(email)
	if emailPresent == true {
		result = false
		return interfaces.User{}, result
	}

	userID := uuid.Must(uuid.NewRandom())
	safePassword := helpers.BlackList(password)
	if safePassword != password {
		result = false
		return interfaces.User{}, result
	}
	passwordHash := helpers.HashAndSalt([]byte(safePassword))
	user := interfaces.User{UserID: userID, Username: username, Email: email, Password: passwordHash}
	db := ConnectDB()
	if err := db.Create(&user).Error; err != nil {
		result = false
		return interfaces.User{}, result
	}
	return user, true
}

func UpdateEmail(newEmail string, email string) {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	user.Email = newEmail
	db.Save(user)

}

func CheckPass(email string, password string) bool {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	passCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	switch passCheck {
	case bcrypt.ErrMismatchedHashAndPassword:
		return false
	case bcrypt.ErrHashTooShort:
		return false
	default:
		return true
	}
}

func DeleteUser(email string) {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	db.Delete(user)
}

// Only use durring testing for tear own
func UnscopedDeleteUser(email string) {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	db.Unscoped().Delete(user)
}
