package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age: ", age)

	// every variable declared must be used, otherwise the code will not compile
	var name = "Dan"

	fmt.Println("Name: ", name)

	// short declaration only works in block scope.
	s := "Learning golang!"
	fmt.Println(s)

	/*
		only new varibale can be created on the using :=
		name := "akash"
		this will give error as name var is already defined.

	*/

}
