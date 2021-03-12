package transaction

import (
	"bank-app/database"
	"bank-app/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateTransaction(payeeEmail string, payorEmail string, amount int) bool {
	payee, payor := FindPayeeAndPayor(payeeEmail, payorEmail)
	if !checkUserBalance(payorEmail, amount) {
		return false
	}
	subtractBalance(payor, amount)
	addBalance(payee, amount)

	return true
}

func subtractBalance(payor interfaces.User, amount int) {
	payor.Balence = payor.Account.Balance - amount
	setBalance(payor.Email, payor.Account.Balance)
}

func addBalance(payee interfaces.User, amount int) {
	payee.Account.Balance = payee.Account.Balance + amount
	setBalance(payee.Email, payee.Balence)
}

func TopUpBalance(email string, amount int) {
	database.TopUpAccountBalance(email, amount)
}

func setBalance(email string, newBalance int) {
	database.SetBalance(email, newBalance)
}

func getAccountBalance(email string) int {
	balance := database.GetUserBalance(email)

	return balance
}

func checkUserBalance(email string, amount int) bool {
	balance := database.GetUserBalance(email)
	if balance < amount {
		return false
	}
	return true
}

func FindPayeeAndPayor(payeeEmail string, payorEmail string) (interfaces.User, interfaces.User) {
	payee, payor := database.GetPayeeAndPayor(payeeEmail, payorEmail)

	return payee, payor
}
