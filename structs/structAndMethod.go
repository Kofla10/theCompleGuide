package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your firts name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthday: ")

	var appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthday:  userBirthdate,
		createdAt: time.Now(),
	}
	//pass parameter
	// outputUserDetails(&appUser)
	//method

	for {
		var number int
		fmt.Println("\nMenu ption")
		fmt.Println("1. View user information")
		fmt.Println("2. Clear First Name")
		fmt.Println("3. Clear Last Name")
		fmt.Println("4. Clear Birthday")
		fmt.Println("5. Exit")
		fmt.Print("\n Enter your option => ")

		fmt.Scan(&number)

		switch number {
		case 1:
			appUser.outputUserDetails()
		case 2:
			appUser.clearUserName()
		case 5:
			fmt.Println("Goodbye")
			return
		}

	}

}

// access to that user Struct
// func  outputUserDetails(u *user) { //parameter
func (u user) outputUserDetails() { //method
	fmt.Println("\n\nThe Name is =>", u.firstName, "\nLastname    =>", u.lastName, "\nBirthday    => "+u.birthday, "\nCreated At  =>", u.createdAt, "\n")
}

func (u user) clearUserName() {
	u.firstName = ""
	u.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
