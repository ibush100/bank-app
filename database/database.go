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
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	user.Account.Balance = newBalance

	db.Save(user)
}

func CreateAccount(email string) interfaces.Account {

	userID := validateUserID

	account := interfaces.Account{OwnerID: userID}
	db := ConnectDB()
	db.Create(&account)
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
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	startBalance := user.Account.Balance
	topUpBalance := startBalance + amount
	user.Account.Balance = topUpBalance

	db.Save(user)
}

func FindUser(email string) uuid.UUID {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)

	return user.UserID
}

func GetUserBalance(email string) int {
	db := ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)

	return user.Account.Balance
}

func IsUserPresent(email string) bool {
	userResult := FindUser(email)
	if userResult != uuid.Nil {
		return false
	}
	return true
}

func CreateUser(username string, email string, password string) (interfaces.User, bool) {
	var result bool
	userExists := IsUserPresent(email)
	if userExists == true {
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
	db.Create(&user)
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
