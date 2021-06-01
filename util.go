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

func checkStringsIsEmpty(strs ...string) bool {
	for _, str := range strs {
		if checkStringIsEmpty(str) {
			return true
		}
	}
	return false
}

func checkStringIsEmpty(str string) bool {
	return len(str) == 0
}
