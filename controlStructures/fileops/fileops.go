package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// created function to read value of the balance
func ReadValueBalance(nameFile string) (valueBalance float64, er error) {
	er = nil
	value, err := os.ReadFile(nameFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(value)
	valueBalance, err = strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 500, errors.New("Failed to parse stored balance values.")
	}
	return
}

// I created a function for write the change to balance
func WriteBalanceToFile(balance float64, nameFile string) {
	//we change to type of balance variable
	balanceText := fmt.Sprint(balance)

	os.WriteFile(nameFile, []byte(balanceText), 0644)
}
