package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

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

// generics in structs
type CustomData interface {
	constraints.Ordered | []string
}

type User[T CustomData] struct {
	id   int
	name string
	data T
}

type CustomMap[T comparable, V int | string] map[T]V

// comparable are types which can be compared using == operator

func main() {
	// fmt.Println(first[int](1, 2))
	fmt.Println(first[float64](1.0, 2))

	//type inference
	fmt.Println(first(1.0, 2.0))

	// ~sub types of type

	// user struct
	u := User[float64]{
		id:   0,
		name: "akash",
		data: 12.2,
	}

	fmt.Println("user: ", u)

	d := CustomMap[string, int]{}

	d["akash"] = 3
}
