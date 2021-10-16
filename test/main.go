package main

import "fmt"

type User struct {
	Name   string
	Class  string
	Age    int
	isMale bool
}

type Operation func(u *User)

func NewUser(opts ...Operation) *User {
	user := &User{}
	for _, opt := range opts {
		opt(user)
	}

	return user
}

func SetName(name string) Operation {
	return func(u *User) {
		u.Name = name
	}
}

func SetAge(age int) Operation {
	return func(u *User) {
		u.Age = age
	}
}

func main() {
	user := NewUser(SetAge(18), SetName("cola"))
	fmt.Printf("user:%#v", user)
}
