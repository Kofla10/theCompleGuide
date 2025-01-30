package main

import (
	"fmt"
)

func main() {
	outPrint("Application Calculator\n")

	renueve := generateInput("Renueve: ")
	expense := generateInput("Expense: ")
	taxRate := generateInput("Tax Rate: ")

	ebt, profit, ratio := calculator(renueve, expense, taxRate)

	fmt.Printf("The values are:\nebt => %v\nprofit=> %v\nratio=> %.3v\n", ebt, profit, ratio)

}

func outPrint(print string) {
	fmt.Print(print)
}

// we created input
func generateInput(input string) float64 {
	var value float64
	fmt.Print(input)
	fmt.Scan(&value)
	return value
}

func calculator(renueve float64, expense float64, taxRate float64) (ebt float64, profit float64, ratio float64) {

	ebt = renueve - expense
	profit = ebt * (1 - taxRate/100)
	ratio = (ebt / profit)
	return
}
