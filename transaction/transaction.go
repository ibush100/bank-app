package transaction

import (
	"bank-app/helpers"
	"bank-app/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateTransaction(payeeEmail string, payorEmail string, amount int) {
	payee, payor := FindPayeeAndPayor(payeeEmail, payorEmail)
	subtractBalance(payor, amount)
	addBalance(payee, amount)
}

func subtractBalance(payor interfaces.User, amount int) {
	payor.Balence = payor.Balence - amount
	setBalance(payor.Email, payor.Balence)
}

func addBalance(payee interfaces.User, amount int) {
	payee.Balence = payee.Balence + amount
	setBalance(payee.Email, payee.Balence)
}

func TopUpBalance(email string, amount int) {
	db := helpers.ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	startBalance := user.Balence
	topUpBalance := startBalance + amount
	user.Balence = topUpBalance

	db.Save(user)
}

func setBalance(email string, newBalance int) {
	db := helpers.ConnectDB()
	var user interfaces.User
	db.Where("email = ?", email).First(&user)
	user.Balence = newBalance

	db.Save(user)
}

func FindPayeeAndPayor(payeeEmail string, payorEmail string) (interfaces.User, interfaces.User) {
	db := helpers.ConnectDB()
	var payee interfaces.User
	var payor interfaces.User
	db.Where("email = ?", payeeEmail).First(&payee)
	db.Where("email = ?", payorEmail).First(&payor)

	return payee, payor
}
