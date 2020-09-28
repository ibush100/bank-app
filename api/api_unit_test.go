package api

import (
	"bank-app/database"
	"bank-app/interfaces"
	"bank-app/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	user := interfaces.User{Email: "email@example.com", Username: "User", Password: "password123"}
	addedUser, result := users.CreateUser(user.Username, user.Email, user.Password)
	assert.EqualValuesf(addedUser.Username, user.Username, "Username should be returned")
	assert.EqualValuesf(addedUser.Email, user.Email, "User email should be returned")
	//rework this since we aren't plain texting passwords
	//assert.EqualValuesf(addedUser.Password, user.Password, "User password should be returned")
	assert.NotNil(addedUser.UserID)
	assert.True(result)
}

func TestFindUser(t *testing.T) {
	assert := assert.New(t)
	result := database.FindUser("email@example.com")
	assert.True(result > 1)
}

func TestPrepareToken(t *testing.T) {
	assert := assert.New(t)
	var id uint = 5124521
	result := PrepareToken(id)
	assert.NotNil(result)
}

func TestCheckPass(t *testing.T) {
	assert := assert.New(t)
	password := "password123"
	email := "email@example.com"
	result := checkPass(password, email)
	assert.True(result)
}
