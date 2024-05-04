package main

import "fmt"

/**
	- Type parameter
	- Type inference
	- Type set

**/

// type sets
type minTypes interface {
	~float64 | int
}

func first[T ~float64](a T, b T) T {
	return a
}

func second[T minTypes](a T, b T) T {
	return b
}

func main() {
	// fmt.Println(first[int](1, 2))
	fmt.Println(first[float64](1.0, 2))

	//type inference
	fmt.Println(first(1.0, 2.0))

	// ~sub types of type
}
