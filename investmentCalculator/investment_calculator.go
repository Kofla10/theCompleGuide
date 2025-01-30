package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64 = 1000
	expectedReturRate := 5.5
	var years float64
	const inflationRate = 6.5

	// fmt.Print("Investment Amount: ")
	outputText("investmet Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturRate)

	futureValue, futureRealValue := calculatorFutureValue(investmentAmount, expectedReturRate, years, inflationRate)
	// futureValue := calculatorFutureValue
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//I don't like this form because more process
	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Real Future Value: %.2f\n", futureRealValue)
	// fmt.Printf("Future Value: %v this is a test \n", futureValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future value(adjusted for Infation):", futureRealValue)
	//we can show decimal that we want with %.2f show two decimal
	fmt.Printf("Future Value: %.0f\nFuture Value (adjusted for Infetion): %.2f\n", futureValue, futureRealValue)

	// fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

// func calculatorFutureValue(investmentAmount float64, expectedReturRate float64, years float64, inflationRate float64)  (float64, float64) {
// we defined the name of variables of return
func calculatorFutureValue(investmentAmount float64, expectedReturRate float64, years float64, inflationRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	//if we defined the names and types, we can do the return this shape
	return
}
