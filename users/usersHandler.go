package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"bank-app/interfaces"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var result bool
	body := helpers.ReadBody(r)
	var fomattedUser interfaces.Register
	validate := validator.New()
	err := json.Unmarshal(body, &fomattedUser)
	helpers.HandleErr(err)
	bodyError := validate.Struct(fomattedUser)
	if bodyError != nil {
		result = false
		w.WriteHeader(http.StatusNotFound)
		return
	}

	registerUser, result := CreateUser(fomattedUser.Username, fomattedUser.Email, fomattedUser.Password)
	if result == false {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	helpers.WriteToJson(w, registerUser)

}

func UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	tokenResult := VerifyToken(token)
	if !tokenResult {
		w.WriteHeader(http.StatusForbidden)
		return
	}

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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	body := helpers.ReadBody(r)
	var formattedBody interfaces.User
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)

	if !database.CheckPass(formattedBody.Email, formattedBody.Password) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		token := PrepareToken()
		w.WriteHeader(http.StatusOK)
		//create token
		helpers.WriteToJson(w, token)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("x-token")
	tokenResult := VerifyToken(token)
	if !tokenResult {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	body := helpers.ReadBody(r)
	var formattedBody interfaces.User
	err := json.Unmarshal(body, &formattedBody)
	helpers.HandleErr(err)
	// check pass function
	if database.IsUserPresent(formattedBody.Email) {
		database.DeleteUser(formattedBody.Email)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
