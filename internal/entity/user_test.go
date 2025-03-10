package entity

import "testing"

func TestNewUser(t *testing.T) {
	email := "test@test.com"
	password := "password"
	user, err := NewUser(email, password)
	if err != nil {
		t.Error(err)
	}
	if user.Email != email {
		t.Errorf("Email is not correct")
	}
	if !user.ValidatePassword(password) {
		t.Errorf("Password is not correct")
	}
	if user.Password == password {
		t.Errorf("Password should be hashed")
	}
}
