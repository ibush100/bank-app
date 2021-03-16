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

	if payee.UserID == payor.UserID {
		return false
	}

	if amount < 0 {
		return false
	}
	subtractBalance(payor, amount)
	addBalance(payee, amount)

	return true
}

func subtractBalance(payor interfaces.User, amount int) bool {
	account := database.GetAccount(payor.Email)
	currentAmount := account.Balance
	newBalance := currentAmount - amount
	result := setBalance(payor.Email, newBalance)
	return result
}

func addBalance(payee interfaces.User, amount int) bool {
	account := database.GetAccount(payee.Email)
	currentAmount := account.Balance
	newBalance := currentAmount + amount
	result := setBalance(payee.Email, newBalance)
	return result
}

func TopUpBalance(email string, amount int) bool {
	result := database.TopUpAccountBalance(email, amount)

	return result
}

func setBalance(email string, newBalance int) bool {
	result := database.SetBalance(email, newBalance)

	return result
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
