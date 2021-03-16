package transaction

import (
	"bank-app/database"
	"bank-app/users"
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {

	payee, resultPayee := users.CreateUser("asdf", faker.Email(), "1234")
	payor, resultPayor := users.CreateUser("asdf", faker.Email(), "1234")
	if resultPayee == false {
		os.Exit(1)
	}
	if resultPayor == false {
		os.Exit(1)
	}
	setBalance(payee.Email, 100)
	setBalance(payor.Email, 200)

	CreateTransaction(payee.Email, payor.Email, 100)
	//update because balance isn't being set in user table anymore
	payeeRes, payorRes := FindPayeeAndPayor(payee.Email, payor.Email)
	payeeAccount := database.GetAccount(payeeRes.Email)
	payorAccount := database.GetAccount(payorRes.Email)
	assert.Equal(t, payeeAccount.Balance, 200, "payee that didn't work")
	assert.Equal(t, payorAccount.Balance, 100, "that didn't work")

}
