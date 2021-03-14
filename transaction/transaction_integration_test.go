package transaction

import (
	"bank-app/helpers"
	"bank-app/users"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/stretchr/testify/assert"
)

type transaction struct {
	PayorEmail string
	PayeeEmail string
	Amount     int
}

type UpdateUser struct {
	Username string
	Password string
	Email    string
	TopUp    int
}

func TestCreateTransactionIntegration(t *testing.T) {
	payee, resultPayee := users.CreateUser("asdf", faker.Email(), "1234")
	payor, resultPayor := users.CreateUser("asdf", faker.Email(), "1234")

	if resultPayee == false {
		os.Exit(1)
	}
	if resultPayor == false {
		os.Exit(1)
	}
	//payee, payor := FindPayeeAndPayor("fresh@example.com", "email@example.com")
	setBalance(payee.Email, 100)
	setBalance(payor.Email, 200)
	transaction := transaction{payee.Email, payor.Email, 100}
	requestByte, _ := json.Marshal(transaction)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/transaction", requestReader)
	token := users.PrepareToken()
	req.Header.Set("x-token", token)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTransactionHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code, "that didn't work")
}

func TestUpdateUserBalance(t *testing.T) {
	user, result := users.CreateUser("asdf", faker.Email(), "1234")

	if result == false {
		os.Exit(1)
	}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("PUT", "/updateBalance", requestReader)
	token := users.PrepareToken()
	req.Header.Set("x-token", token)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUserBalance)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code, "that didn't work")
}
