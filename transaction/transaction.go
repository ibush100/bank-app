package transaction

import (
	"bank-app/helpers"
	"bank-app/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateTransaction(payeeEmail string, payorEmail string, amount int) {
	// find users
	payee, payor := findPayeeAndPayor(payeeEmail, payorEmail)

	//create transaction
	subtractBalance(payee, amount)
	addBalance(payor, amount)
	// updateBalance()
}

func subtractBalance(payor interfaces.User, amount int) {
	//update balance here
	payor.Balence = payor.Balence - amount

}

func addBalance(payee interfaces.User, amount int) {
	//update balance here
	payee.Balence = payee.Balence - amount
}

func findPayeeAndPayor(payeeEmail string, payorEmail string) (interfaces.User, interfaces.User) {
	db := helpers.ConnectDB()
	var payee interfaces.User
	var payor interfaces.User
	db.Where("email = ?", payeeEmail).First(&payee)
	db.Where("email = ?", payorEmail).First(&payor)

	return payee, payor
}
