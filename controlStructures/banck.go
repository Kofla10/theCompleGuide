package main

import "fmt"

func main() {

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("2. Withdraw money")
	fmt.Println("Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {

	}

	fmt.Println("Your choice: ", choice)
}
