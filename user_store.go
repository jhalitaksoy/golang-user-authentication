package main

import "net/http"

var userStore UserStore = newUserStore()

//UserStore ...
type UserStore interface {
	Register(user *User) int
	Login(user *User) (int, string)
}

type UserStoreImpl struct {
	Users []*User
}

func newUserStore() *UserStoreImpl {
	return &UserStoreImpl{
		Users: make([]*User, 0),
	}
}

func (userStore *UserStoreImpl) Register(user *User) int {
	for _, eachUser := range userStore.Users {
		if eachUser.Name == user.Name {
			return http.StatusConflict
		}
	}
	userStore.AddUser(*user)
	return http.StatusOK
}

func (userStore *UserStoreImpl) Login(user *User) (int, string) {
	for _, eachUser := range userStore.Users {
		if eachUser.Name == user.Name && eachUser.Password == user.Password {
			return http.StatusOK, eachUser.ID
		}
	}
	return http.StatusConflict, ""
}

func (userStore *UserStoreImpl) AddUser(user User) {
	userStore.Users = append(userStore.Users, &user)
}
