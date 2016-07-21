package main

import (
	"fmt"
	"roastme/db"
	domain "roastme/domain"
)

func main() {
	// set up DB session
	db := mongodb.New()
	db.NewSession()
	user := domain.NewUser("Antonio Diaz", "jadiaz", "1234", "jadiaz@paradigmadigital.com")
	fmt.Println(user)
	err := db.Insert("users", user)

	if err != nil {
		panic(err)
	}
	fmt.Println("Document inserted successfully!")
}
