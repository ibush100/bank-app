package users

import (
	"bank-app/database"
	"bank-app/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"

	"github.com/stretchr/testify/assert"
)

type UpdateEmail struct {
	Email    string
	NewEmail string
	Password string
}

type user struct {
	Email    string
	Username string
	Password string
}

func TestRegisterUser(t *testing.T) {
	user := user{Email: faker.Email(), Username: "asdfsadf", Password: "password"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/user", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code, "that didn't work")
}

func TestRegisterUserBlackListPassword(t *testing.T) {
	user := user{Email: "email@example.com", Username: "User", Password: "passwor*d="}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/user", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code, "that didn't work")
}

func TestLoginUser(t *testing.T) {
	user, result := CreateUser(faker.Name(), faker.Email(), "password123")
	if result == false {
		assert.FailNow(t, "user was not created")
	}
	user.Password = "password123"
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/login", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code, "that didn't work")

	database.UnscopedDeleteUser(user.Email)
}

func TestLoginUserWrongPassword(t *testing.T) {
	user := user{Email: "fh@example.com", Password: "passasdwd123"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/login", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 403, rr.Code, "that didn't work")
}

func TestUpdateUserEmail(t *testing.T) {
	user := UpdateEmail{Email: "new@example.com", NewEmail: "fresh@example.com"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("PUT", "/user", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUserEmail)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code, "that didn't work")
}
