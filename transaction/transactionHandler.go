package transaction

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"bank-app/users"
	"encoding/json"
	"net/http"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	tokenResult := users.VerifyToken(token)
	if !tokenResult {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	body := helpers.ReadBody(r)
	var formattedBody interfaces.Transaction
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if !database.IsUserPresent(formattedBody.PayorEmail) || !database.IsUserPresent(formattedBody.PayeeEmail) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	result := CreateTransaction(formattedBody.PayeeEmail, formattedBody.PayorEmail, formattedBody.Amount)
	if !result {
		w.WriteHeader(http.StatusForbidden)
		return
	}
}

func UpdateUserBalance(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	tokenResult := users.VerifyToken(token)
	if !tokenResult {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	body := helpers.ReadBody(r)
	var formattedBody interfaces.UpdateUserBalance
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if database.IsUserPresent(formattedBody.Email) {
		TopUpBalance(formattedBody.Email, formattedBody.TopUp)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
