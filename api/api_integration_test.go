package api

import (
	"bank-app/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	Email    string
	Username string
	Password string
}

type UpdateEmail struct {
	Email    string
	NewEmail string
	Password string
}

type UpdateBalance struct {
	Email    string
	TopUp    int
	Password string
}

func TestRegisterUser(t *testing.T) {
	user := user{Email: "email@example.com", Username: "User", Password: "password123"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/register", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(registerUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code, "that didn't work")
}

func TestLoginUser(t *testing.T) {
	user := user{Email: "fresh@example.com", Password: "password123"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/login", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 201, rr.Code, "that didn't work")
}

func TestLoginUserWrongPassword(t *testing.T) {
	user := user{Email: "fresh@example.com", Password: "passwd123"}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/login", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(loginUser)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 403, rr.Code, "that didn't work")
}

func TestUpdateUserEmail(t *testing.T) {
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

func TestUpdateUserBalance(t *testing.T) {
	user := UpdateBalance{Email: "fresh@example.com", TopUp: 100}
	requestByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("PUT", "/updateBalance", requestReader)
	if err != nil {
		helpers.HandleErr(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateUserBalance)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code, "that didn't work")
}
