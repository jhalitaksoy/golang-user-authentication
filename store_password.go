package main

import (
	"errors"
)

//PasswordStore ...
type PasswordStore interface {
	Create(userID, hash string) error
	Get(userID string) *string
	Update(userID, hash string)
	Delete(userID string)
}

type PasswordStoreImpl struct {
	Passwords map[string]*string
}

func newPasswordStoreImpl() *PasswordStoreImpl {
	return &PasswordStoreImpl{
		Passwords: make(map[string]*string),
	}
}

func (passwordStore *PasswordStoreImpl) Create(userID, hash string) error {
	passwordStored := passwordStore.Passwords[userID]
	if passwordStored != nil {
		return errors.New("there is a already password")
	}

	passwordStore.Passwords[userID] = &hash

	return nil
}

func (passwordStore *PasswordStoreImpl) Get(userID string) *string {
	return passwordStore.Passwords[userID]
}

func (passwordStore *PasswordStoreImpl) Update(userID, hash string) {
	passwordStore.Passwords[userID] = &hash
}

func (passwordStore *PasswordStoreImpl) Delete(userID string) {
	passwordStore.Passwords[userID] = nil
}
