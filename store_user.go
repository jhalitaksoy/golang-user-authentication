package main

import (
	"errors"
	"fmt"
)

//UserStore ...
type UserStore interface {
	Create(user *User) error
	GeById(id string) *User
	GetByName(name string) *User
	Update(user *User)
	Delete(id string)
}

type UserStoreImpl struct {
	lastUserID int
	Users      map[string]*User
}

func newUserStoreImpl() *UserStoreImpl {
	return &UserStoreImpl{
		lastUserID: -1,
		Users:      make(map[string]*User),
	}
}

func (userStore *UserStoreImpl) Create(user *User) error {
	userStored := userStore.GetByName(user.Name)
	if userStored != nil {
		return errors.New("user name is not suitable")
	}

	id := userStore.createNewUserId()

	user.ID = fmt.Sprint(id)

	userStore.Users[user.ID] = user
	return nil
}

func (userStore *UserStoreImpl) GeById(id string) *User {
	return userStore.Users[id]
}

func (userStore *UserStoreImpl) GetByName(name string) *User {
	for _, user := range userStore.Users {
		if user.Name == name {
			return user
		}
	}

	return nil
}

func (userStore *UserStoreImpl) Update(user *User) {
	userStore.Users[user.ID] = user
}

func (userStore *UserStoreImpl) Delete(id string) {
	userStore.Users[id] = nil
}

func (userStore *UserStoreImpl) createNewUserId() int {
	userStore.lastUserID++
	return userStore.lastUserID
}
