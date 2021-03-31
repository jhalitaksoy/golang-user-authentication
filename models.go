package main

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func newUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
