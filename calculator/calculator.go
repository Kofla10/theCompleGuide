package main

import "fmt"

func main() {

	fmt.Println("Application Calculator")

	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expense: ")
	fmt.Scan(&expense)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	// earnings before tax
	ebt := revenue - expense
	// earnings after tax
	profit := ebt * (1 - taxRate/100)

	ratio := (ebt / profit)

	fmt.Println("EBT:", ebt)

	fmt.Println("Profir:", profit)

	fmt.Println("Ratio: ", ratio)

}
