package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

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
	err := hashAndSaltUserPassword(user)
	if err != nil {
		log.Printf("An error occured while hashing password , error : %v\n", err)
		return http.StatusInternalServerError
	}
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
		if eachUser.Name == user.Name {
			hash := []byte(eachUser.Password)
			plainText := []byte(user.Password)
			err := bcrypt.CompareHashAndPassword(hash, plainText)
			if err != nil {
				return http.StatusConflict, ""
			} else {
				jwt := userStore.CreateJWTToken(eachUser)
				return http.StatusOK, jwt
			}

		}
	}
	return http.StatusConflict, ""
}

func (userStore *UserStoreImpl) CreateJWTToken(user *User) string {
	signedToken, err := CreateJWTToken(user)
	if err != nil {
		panic(err)
	}
	return signedToken
}

func (userStore *UserStoreImpl) AddUser(user User) {
	userStore.Users = append(userStore.Users, &user)
}
