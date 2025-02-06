package main

import (
	"fmt"
	"time"

	"example.com/menu"
	"example.com/writetext"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main() {
	userFirstName := addUserData("Please enter your firts name: ")
	userLastName := addUserData("Please enter your last name: ")
	userBirthdate := addUserData("Please enter your birthday: ")

	var appUser = newUser(userFirstName, userLastName, userBirthdate)

	for {
		var number int
		menu.Menu()

		fmt.Scan(&number)

		switch number {
		case 1:
			appUser.outputUserDetails()
		case 2:
			appUser.edittUser()
		case 3:
			appUser.clearUserName()
		case 4:
			appUser.clearLastName()
		case 5:
			appUser.clearBirthday()
		case 6:
			appUser.clearAll()
		case 7:
			fmt.Println("Goodbye")
			return
		}
	}
}

func newUser(firstName, lastName, birthday string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}

}

// This fuction print the user data and write in the file
func (u *user) outputUserDetails() {

	write := fmt.Sprint("\n\nFistName    => ", u.firstName, "\nLastname    => ", u.lastName, "\nBirthday    => "+u.birthday, "\nCreated At  => ", u.createdAt, "\n")

	writetext.WriteFile(write)
	fmt.Print(write)
}

// Thishis fucntion edit the user data
func (u *user) edittUser() {
	firstName := addUserData("Please enter the new firts name: ")
	lastName := addUserData("Please enter the new last name: ")
	birthday := addUserData("Please enter the new birthday: ")

	u.firstName = firstName
	u.lastName = lastName
	u.birthday = birthday
}

// Thishis function remove the userName of user data
func (u *user) clearUserName() {
	u.firstName = ""
	u.outputUserDetails()
}

// Thishis function remove the lastName of user data
func (u *user) clearLastName() {
	u.lastName = ""
	u.outputUserDetails()
}

// Thishis function remove the birthday of user data
func (u *user) clearBirthday() {
	u.birthday = ""
	u.outputUserDetails()
}

// Thishis fucntion remove the user data
func (u *user) clearAll() {
	u.firstName = ""
	u.lastName = ""
	u.birthday = ""
	u.outputUserDetails()
}

// This function add the value to user data
func addUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
