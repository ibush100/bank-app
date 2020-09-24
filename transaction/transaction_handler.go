package transaction

import (
	"bank-app/helpers"
	"bank-app/interfaces"
	"bank-app/users"
	"encoding/json"
	"net/http"
)

func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody interfaces.Transaction
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if users.IsUserPresent(formattedBody.PayorEmail) && users.IsUserPresent(formattedBody.PayeeEmail) {
		CreateTransaction(formattedBody.PayeeEmail, formattedBody.PayorEmail, formattedBody.Amount)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
