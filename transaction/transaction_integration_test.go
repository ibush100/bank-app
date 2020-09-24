package transaction

import (
	"bank-app/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type transaction struct {
	payorEmail string
	PayeeEmail int
	Password   string
}

//Need to work on this
func TestCreateTransactionIntegration(t *testing.T) {

	CreateTransaction("fresh@example.com", "email@example.com", 100)
	payee, payor := FindPayeeAndPayor("fresh@example.com", "email@example.com")
	assert.Equal(t, payee.Balence, 200, "that didn't work")
	assert.Equal(t, payor.Balence, 100, "that didn't work")

	user := UpdateEmail{Email: "new@example.com", NewEmail: "fresh@example.com"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("PUT", "/updateEmail", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateUserEmail)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code, "that didn't work")

}
