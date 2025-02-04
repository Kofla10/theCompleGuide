package main

import (
	"fmt"

	"exercise.go/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, problem := fileops.ReadValueBalance(accountBalanceFile)
	if problem != nil {
		fmt.Printf("\n\n---------------------------------------\nWe have error when I try reading file, the error is: %s\n---------------------------------------\n\n", problem)
		// panic("Can't con tinue, sorry")
	}
	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOption()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much do you want to deposit?: ")
			fmt.Scan(&depositAmount)
			if depositAmount > 0 {
				accountBalance += depositAmount
				fmt.Printf("The amount deposit is: %v\n.\n..\n...\ndone\n", depositAmount)
			} else {
				fmt.Println("You can't continue because the deposit negative")
				continue
			}
		case 3:
			var yesno int
			fmt.Print("Are you sure that you want to withdraw you money?\n1. yes\n2.no\n")
			fmt.Scan(&yesno)
			if yesno == 1 {
				var amount float64
				fmt.Printf("What is the amount that you want to withdraw? ")
				fmt.Scan(&amount)
				if accountBalance >= amount && amount > 0 {
					accountBalance -= amount
				} else {
					fmt.Println("insufficient found")
					continue
				}
			} else {
				fmt.Println("Your request was cancelled")
				continue
			}

		default:
			fmt.Println("GoodBye.")
			return
			// don't use break in go, we use return
			// break
		}

		fileops.WriteBalanceToFile(accountBalance, accountBalanceFile)
		fmt.Printf("\nTha amount of money that you have in you acount is: %v\n\n", accountBalance)
	}
}
