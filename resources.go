package main

type Resource interface{}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}
