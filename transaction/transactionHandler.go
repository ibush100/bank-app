package transaction

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"encoding/json"
	"net/http"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody interfaces.Transaction
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if database.IsUserPresent(formattedBody.PayorEmail) && database.IsUserPresent(formattedBody.PayeeEmail) {
		CreateTransaction(formattedBody.PayeeEmail, formattedBody.PayorEmail, formattedBody.Amount)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func UpdateUserBalance(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody interfaces.UpdateUserBalance
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if database.IsUserPresent(formattedBody.Email) {
		transaction.TopUpBalance(formattedBody.Email, formattedBody.TopUp)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
