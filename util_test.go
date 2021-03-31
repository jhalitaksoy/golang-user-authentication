package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashAndSaltUserPassword(t *testing.T) {
	user := newUser("hlt", "1234")
	user2 := *user

	hashAndSaltUserPassword(user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user2.Password))
	if err != nil {
		t.Fail()
	}
}
