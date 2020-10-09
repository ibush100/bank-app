package transaction

import (
	"bank-app/database"
	"bank-app/helpers"
	"encoding/json"
	"net/http"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody Transaction
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if !database.IsUserPresent(formattedBody.PayorEmail) || !database.IsUserPresent(formattedBody.PayeeEmail) {
		w.WriteHeader(http.StatusForbidden)
	}
	result := CreateTransaction(formattedBody.PayeeEmail, formattedBody.PayorEmail, formattedBody.Amount)
	if !result {
		w.WriteHeader(http.StatusForbidden)
	}
}

func UpdateUserBalance(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody UpdateUserBalance
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if database.IsUserPresent(formattedBody.Email) {
		TopUpBalance(formattedBody.Email, formattedBody.TopUp)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
