package transaction

import (
	"bank-app/helpers"
	"bank-app/users"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type transaction struct {
	PayorEmail string
	PayeeEmail string
	Amount     int
}

func TestCreateTransactionIntegration(t *testing.T) {

	payee, payor := FindPayeeAndPayor("fresh@example.com", "email@example.com")
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
