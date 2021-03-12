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
	account := database.GetAccount(payor.Email)
	currentAmount := account.Balance
	newBalance := currentAmount - amount
	setBalance(payor.Email, newBalance)
}

func addBalance(payee interfaces.User, amount int) {
	account := database.GetAccount(payee.Email)
	currentAmount := account.Balance
	newBalance := currentAmount - amount
	setBalance(payee.Email, newBalance)
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
