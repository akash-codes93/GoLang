package main

import "fmt"

/* scopes are of three types */
// 1. package scope
// 2. block scope
// 3. file scope

var numbers [10]int

func init() {
	fmt.Println("This is init #1")
}

func init() {
	fmt.Println("This is init #2") // this is allowed

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

func main() {
	pl := fmt.Println
	sayHello("Akash")

	tf := toFahrenheit(boilingPoint)
	pl(tf)

	pl(numbers)
}

// go run *.go :: This is how we run multiple files in the same directory
// imports are file scope..that means that fmt must be imported in hello.go

// const function and variable are package scope :: that is why sayHello is available to main.go
// same name cannot be declared in the same scope that means sayHello cannot be declared in main.go

// init() is called automatically at the start, you cannot call init() explicitly.
// init is used to initialize global variable which cannot be declared otherwise.
