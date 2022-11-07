package main

type User struct {
	Name string
}

var interfaceUser UserInterface = &User{}

type UserInterface interface {
	GetUserName() string
}

func (u *User) GetUserName() string {
	return u.Name
}
