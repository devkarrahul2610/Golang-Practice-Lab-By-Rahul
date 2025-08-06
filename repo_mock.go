package main

import "fmt"

type MockRepository struct{}

func (m MockRepository) GetUserByID(id int) (User, error) {
	fmt.Println("Fetching user from Mock")

	return User{Id: 1, Name: "shubham"}, nil
}
