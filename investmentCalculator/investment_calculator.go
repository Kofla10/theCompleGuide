package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var expectdReturRate = 5.5
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1+expectdReturRate/100, years)

	fmt.Println(futureValue)
}
