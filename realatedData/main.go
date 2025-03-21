package main

import "fmt"

func main() {
	productName := [4]string{}

	productName = [4]string{"Camilo"}
	fmt.Println(productName)
	test := []any{"Camilo", "prueba 1", 2, 3.2 + 3, "asdfasd", 123, 23, 12, 3, 23, 243, 23}
	prices := [5]float64{12.2, 423.2, 23.0, 34.5, 123.1}
	fmt.Println(prices)
	fmt.Println(test)

	array()
}

func array() {
	fmt.Print("\n\n\n")
	test := [4]int{1, 2, 4, 3}
	fmt.Println(test)

	// we created another array to add values of the last array
	test2 := test[1:4]

	fmt.Println(test2)
}
