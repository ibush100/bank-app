package transaction

import (
	"bank-app/helpers"
	"bank-app/interfaces"
)

func GetToBalance(userID uint) int {
	db := helpers.ConnectDB()
	var user interfaces.User
	db.Where("user_ID = ?", userID).First(&user)

	return user.Balence
}
