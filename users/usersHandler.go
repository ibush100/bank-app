package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"encoding/json"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)

	var fomattedUser interfaces.Register
	err := json.Unmarshal(body, &fomattedUser)
	helpers.HandleErr(err)
	registerUser, result := CreateUser(fomattedUser.Username, fomattedUser.Email, fomattedUser.Password)
	if result {
		w.WriteHeader(http.StatusCreated)
		helpers.WriteToJson(w, registerUser)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody interfaces.UpdateUserEmail
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if database.IsUserPresent(formattedBody.Email) {
		database.UpdateEmail(formattedBody.NewEmail, formattedBody.Email)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

// this will break until you put bcryt stuff in
func LoginUser(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody interfaces.User
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	if checkPass(formattedBody.Email, formattedBody.Password) {
		token := PrepareToken(formattedBody.ID)
		w.WriteHeader(http.StatusCreated)
		// create token
		helpers.WriteToJson(w, token)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
