package main

import (
	"fmt"

	"example.com/menu"
	"example.com/user"
)

func main() {

	var appUser *user.User

	userFirstName, userLastName, userBirthdate := fieldValue()
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	admin := user.NewAdmin("juankofla@gmail.com", "1234")

	admin.OutputDetails()

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputDetails()
	for {
		var number int
		menu.Menu()

		fmt.Scan(&number)

		switch number {
		case 1:
			appUser.OutputDetails()
		case 2:
			editFirstName, editLastName, editBirthday := fieldValue()
			appUser.Edit(editFirstName, editLastName, editBirthday)
		case 3:
			appUser.RemoveName()
		case 4:
			appUser.RemoveLastName()
		case 5:
			appUser.RemoveBirthday()
		case 6:
			appUser.Remove()
		case 7:
			fmt.Println("Goodbye")
			return
		}
	}
}

// This function add the value to user data
func addUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func fieldValue() (name string, lastName string, birthday string) {

	name = addUserData("Please enter your firts name: ")
	lastName = addUserData("Please enter your last name: ")
	birthday = addUserData("Please enter your birthday: ")
	return
}
