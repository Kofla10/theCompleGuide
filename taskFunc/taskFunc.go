package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var nameFile = "value.txt"

func main() {
	outPrint("Application Calculator\n")

	renueve, err := generateInput("Renueve: ")
	if err != nil {
		printError(err, "renueve")
	}

	expense, err := generateInput("Expense: ")
	if err != nil {
		printError(err, "expense")
	}

	taxRate, err := generateInput("Tax Rate: ")
	if err != nil {
		printError(err, "taxRate")
	}

	ebt, profit, ratio := calculator(renueve, expense, taxRate)

	writeFile(ebt, profit, ratio)
	fmt.Printf("The values are:\nebt    => %v\nprofit => %v\nratio  => %.3v\n", ebt, profit, ratio)

}

func printError(print error, field string) {
	fmt.Println("\n", field, "=>", print, "\n")
	panic("Can't continuer")
}

func writeFile(ebt float64, ratio float64, profit float64) {
	// data := fmt.Sprint("The values are:\nebt    => %.3f\nprofit => %.3f\nratio  => %.3f", ebt, profit, ratio)
	data := fmt.Sprint("The values are:\nebt    => " + strconv.FormatFloat(ebt, 'f', 2, 64) + "\nprofit => " + strconv.FormatFloat(profit, 'f', 2, 64) + "\nratio  => " + strconv.FormatFloat(ratio, 'f', 2, 64) + "\n")
	os.WriteFile(nameFile, []byte(data), 0644)
}

func outPrint(print string) {
	fmt.Print(print)
}

// we created input
func generateInput(input string) (value float64, err error) {
	err = nil
	fmt.Print(input)
	fmt.Scan(&value)

	if value == 0 {
		err = errors.New("Error: The values is 0")
	} else if value < 0 {
		err = errors.New("Error: The value is < 0")
	}

	return value, err
}

func calculator(renueve float64, expense float64, taxRate float64) (ebt float64, profit float64, ratio float64) {

	ebt = renueve - expense
	profit = ebt * (1 - taxRate/100)
	ratio = (ebt / profit)
	return
}
