package api

import (
	"bank-app/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	user := interfaces.User{Email: "email@example.com", Username: "User", Password: "password123"}
	addedUser, result := createUser(user.Username, user.Email, user.Password)
	assert.EqualValuesf(addedUser.Username, user.Username, "Username should be returned")
	assert.EqualValuesf(addedUser.Email, user.Email, "User email should be returned")
	assert.EqualValuesf(addedUser.Password, user.Password, "User password should be returned")
	assert.NotNil(addedUser.UserID)
	assert.True(result)
}

func TestFindUser(t *testing.T) {
	assert := assert.New(t)
	result := FindUser("email@example.com")
	assert.True(result > 1)
}

func TestPrepareToken(t *testing.T) {
	assert := assert.New(t)
	var id uint = 5124521
	result := PrepareToken(id)
	assert.NotNil(result)
}
