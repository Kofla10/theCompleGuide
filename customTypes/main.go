package main

import "fmt"

// we can create diferent types with diferent names and with its methods
// we create the type str is same to string.
type str string

//we created a method of type str
func (s str) log(){
	fmt.Println(s)
}

func main() {
	var name str = "She is a person famous who has a lot of money that buy much car"
	str.log(name)
}