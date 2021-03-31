package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashAndSaltUserPassword(user *User) error {
	bytes, error := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if error != nil {
		return error
	}
	user.Password = string(bytes)
	return nil
}
