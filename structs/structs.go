package main

import (
	"fmt"
	"time"
)

type user1 struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main1() {
	userFirstName := getUserData1("Please enter your firts name: ")
	userLastName := getUserData1("Please enter your last name: ")
	userBirthdate := getUserData1("Please enter your birthday: ")

	var appUser1 = user1{
		firstName: userFirstName,
		lastName:  userLastName,
		birthday:  userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetails1(appUser1)
	//... do something awesome with that gathered data!

	// fmt.Printf("\n\nFistName   => %s\n", appUser.firstName)
	// fmt.Printf("LastName   => %s\n",  appUser.lastName)
	// fmt.Printf("Birthday   => %s\n",  appUser.birthday)
	// fmt.Printf("Created At => %v\n\n", appUser.createdAt)

}

func outputUserDetails1(u user1) {
	fmt.Println(u.firstName)
	fmt.Println(u.lastName)
	fmt.Println(u.birthday)
	fmt.Println(u.createdAt)
}

func getUserData1(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
