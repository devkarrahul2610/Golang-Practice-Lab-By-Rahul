package main

import "fmt"

type Postgres struct{}

func (p Postgres) GetUserByID(Id int) (User, error) {
	fmt.Println("Fetching user from Postgres database")

	return User{Id: 1, Name: "shubham"}, nil
}
