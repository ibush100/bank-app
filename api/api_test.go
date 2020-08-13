package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	user := User{Email: "email@example.com", Username: "User", Password: "password123"}
	addedUser := createUser(user.Username, user.Email, user.Password)
	assert.Equal(addedUser, user, "User should be returned")
}
