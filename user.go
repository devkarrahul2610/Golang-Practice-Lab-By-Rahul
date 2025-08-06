package main

type User struct {
	Id   int
	Name string
}

type UserRepository interface {
	GetUserByID(id int) (User, error)
}
