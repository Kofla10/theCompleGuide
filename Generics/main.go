// [T int | float64 | string]: "This function works with a type T, and T can be either int, float64, or string." This is the key part that makes the function generic. It means the function can handle different types without needing separate implementations.
package main

import "fmt"

func main() {
	fmt.Println("Generics")

	result := add(1, 1)

	fmt.Println(result)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
