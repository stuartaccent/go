package main

import (
	"app/conf"
	"app/models"
	"flag"
	"fmt"
	"log"
)

// create a new user
func main() {
	email := flag.String("email", "", "email")
	firstname := flag.String("firstname", "", "firstname")
	lastname := flag.String("lastname", "", "lastname")
	password := flag.String("password", "", "password")

	flag.Parse()

	db, err := conf.ConnectDB()
	if err != nil {
		log.Println(err)
		return
	}

	user := &models.User{
		Email:     *email,
		FirstName: *firstname,
		LastName:  *lastname,
	}
	user.SetPassword(*password)
	if err := db.Create(user).Error; err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Successfully created user")
}
