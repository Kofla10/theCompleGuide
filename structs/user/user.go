package user

import (
	"errors"
	"fmt"
	"time"

	"example.com/writetext"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func New(firstName, lastName, birthday string) (*User, error) {
	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("**ERROR:** All field are required\n\n\nThis application END\n\n")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil

}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Camilo",
			lastName:  "Peñaranda",
			birthday:  "--/--/--",
			createdAt: time.Now(),
		},
	}
}

// This fuction print the User data and write in the file
func (u *User) OutputDetails() {

	write := fmt.Sprint("\n\nFistName    => ", u.firstName, "\nLastname    => ", u.lastName, "\nBirthday    => "+u.birthday, "\nCreated At  => ", u.createdAt, "\n")

	writetext.WriteFile(write)
	fmt.Print(write)
}

// Thishis fucntion edit the User data
func (u *User) Edit(firstName, lastName, birthday string) {
	u.firstName = firstName
	u.lastName = lastName
	u.birthday = birthday

	u.OutputDetails()
}

// Thishis function remove the UserName of User data
func (u *User) RemoveName() {
	u.firstName = ("ease enter the new firts name: ")
	u.OutputDetails()
}

// Thishis function remove the lastName of User data
func (u *User) RemoveLastName() {
	u.lastName = ""
	u.OutputDetails()
}

// Thishis function remove the birthday of User data
func (u *User) RemoveBirthday() {
	u.birthday = ""
	u.OutputDetails()
}

// Thishis fucntion remove the User data
func (u *User) Remove() {
	u.firstName = ""
	u.lastName = ""
	u.birthday = ""
	u.OutputDetails()
}
