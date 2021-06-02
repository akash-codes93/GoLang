package main

import "fmt"

func main() {

	var i1 int8 = 100 // its range is from -128 to +127 (2^7)
	// var i1 int8 = -129 // this will give compile time error
	fmt.Printf("%T\n", i1)

	var i rune = 12345 // rune is alias for int32
	fmt.Printf("%T\n", i)

	var ip rune = 'f'
	fmt.Printf("%d\n", ip) // print 108 which is decimal code for f
	fmt.Printf("%x\n", ip) // hexa decimal value for ascii code

	// STRING
	var str string = "akash"
	fmt.Println(str)

	// array
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// slice : a slice can have dynamic length
	var cities = []string{"india", "pakistan", "tokyo", "zimbabwe"}
	fmt.Printf("%T\n", cities)

	// map :  similar to dict in python
	var balances = map[string]float64{
		"USD": 60.9,
		"INR": 1,
		"PAK": 23, // comma needed at the end
	}

	fmt.Printf("%T\n", balances)
	fmt.Println(balances)

	// struct type : similar to class type
	type Person struct {
		name string
		age  int8
	}

	var akash Person
	fmt.Printf("%T\n", akash)

	// pointer
	var x int = 2
	ptr := &x

	fmt.Printf("%T\n", ptr)
	fmt.Println(ptr)

	// function type
	fmt.Printf("%T\n", emptyFunction)

}

func emptyFunction() {

}
