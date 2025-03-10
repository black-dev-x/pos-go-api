package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	email := "test@test.com"
	password := "password"
	user, err := NewUser(email, password)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.True(t, user.ValidatePassword(password))
	assert.NotEqual(t, password, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	email := "t@t.com"
	password := "password"
	user, _ := NewUser(email, password)
	assert.True(t, user.ValidatePassword(password))
	assert.False(t, user.ValidatePassword("wrong"))
	assert.NotEqual(t, password, user.Password)
}
