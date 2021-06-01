package main

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}

type LoginParameters struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterParameters struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
