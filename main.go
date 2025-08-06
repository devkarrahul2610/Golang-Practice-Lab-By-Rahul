package main

import (
	"fmt"
	"log"
)

func main() {
	repo := Postgres{}
	service := UserService{repo: repo}
	user, err := service.FetchUserData(1)
	if err != nil {

		log.Fatalf("Error fetching user: %v", err)
		return
	}

	fmt.Println("User fetched successfully:", user)
}
