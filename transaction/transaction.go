package transaction

import (
	"bank-app/database"
	"bank-app/users"

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

func subtractBalance(payor users.User, amount int) {
	payor.Balence = payor.Balence - amount
	setBalance(payor.Email, payor.Balence)
}

func addBalance(payee users.User, amount int) {
	payee.Balence = payee.Balence + amount
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

func FindPayeeAndPayor(payeeEmail string, payorEmail string) (users.User, users.User) {
	payee, payor := database.GetPayeeAndPayor(payeeEmail, payorEmail)

	return payee, payor
}
