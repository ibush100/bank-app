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

func TestRegisterUser(t *testing.T) {
	user := user{Email: "fresh@example.com", Username: "User", Password: "password123"}
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
