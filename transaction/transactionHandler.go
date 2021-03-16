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
	if !database.IsUserPresent(formattedBody.PayorEmail) || !database.IsUserPresent(formattedBody.PayeeEmail) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	result := CreateTransaction(formattedBody.PayeeEmail, formattedBody.PayorEmail, formattedBody.Amount)
	if !result {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusCreated)
	helpers.WriteToJson(w, "transaction succesful")
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

	userPresent := database.IsUserPresent(formattedBody.Email)

	if userPresent == false {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	topUpResult := TopUpBalance(formattedBody.Email, formattedBody.TopUp)
	if topUpResult == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	helpers.WriteToJson(w, "topUp Sucessful")
}
