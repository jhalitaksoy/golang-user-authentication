package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashAndSaltPassword(password string) (string, error) {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if error != nil {
		return "", error
	}
	hash := string(bytes)
	return hash, nil
}
