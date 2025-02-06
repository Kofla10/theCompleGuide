package main

import (
	"fmt"
	"time"
)

type user2 struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main2() {
	userFirstName := getUserData2("Please enter your firts name: ")
	userLastName := getUserData2("Please enter your last name: ")
	userBirthdate := getUserData2("Please enter your birthday: ")

	var appUser = user2{
		firstName: userFirstName,
		lastName:  userLastName,
		birthday:  userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)

}

func outputUserDetails(u *user2) {
	fmt.Println("\n\nThe Name is =>"+u.firstName, "\nLastname =>"+u.lastName, "\nBirthday =>"+u.birthday+"\nCreated At", u.createdAt, "\n")
}

func getUserData2(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
