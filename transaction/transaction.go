package transaction

import (
	"bank-app/database"
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
	database.TopUpAccountBalance(email, amount)
}

func setBalance(email string, newBalance int) {
	database.SetBalance(email, newBalance)
}

func FindPayeeAndPayor(payeeEmail string, payorEmail string) (interfaces.User, interfaces.User) {
	payee, payor := database.GetPayeeAndPayor(payeeEmail, payorEmail)

	return payee, payor
}
