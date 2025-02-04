package main

import "fmt"

func main() {
	age := 32 //regula variable

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	adultYear := getAdultYear(agePointer)
	fmt.Println(adultYear)

	fmt.Println(*agePointer)
	fmt.Println("Age without pointer", age)
}

func getAdultYear(age *int) int {
	*age = *age - 18
	return *age
}
