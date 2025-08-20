package main

import "log"

func FetchUserDataTest(Id int) (User, error) {
	repo := MockRepository{}
	userService := UserService{repo: repo}
	user, err := userService.FetchUserData(1)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	return user, nil
}
