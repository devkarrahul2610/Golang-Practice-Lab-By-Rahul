package main

import (
	"log"
)

type UserService struct {
	repo UserRepository
}

func (r UserService) FetchUserData(Id int) (User, error) {
	User, err := r.repo.GetUserByID(Id)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	return User, nil
}
