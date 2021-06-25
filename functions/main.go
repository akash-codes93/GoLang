package main

import (
	"fmt"
	"math"
	"strings"
)

// p := fmt.Printf
// pl := fmt.Println

func myFirstFunction() {
	fmt.Printf("This is my first function\n")
}

func functionWithArguments(a int, b int) {
	fmt.Println(a + b)
}

func functionWithArguments2(a, b, c int, d float64, s string) {
	fmt.Println(a, b, c, d, s)
}

func f3(a float64) float64 {
	// function with return
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	// multiple returns
	return a + b, a * b
}

func sum(a, b int) (s int) {
	// s is a named return parameter
	fmt.Println(s) // s -> 0

	s = a + b

	return // by default s is returned this is called naked return
}

// variadic function
func f1(a ...int) { // taking slice as function variable
	fmt.Println(a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.0
	product := 1.0

	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

// mixing variadic and non-variadic method
func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, "")

	returnString := fmt.Sprintf("Age: %d, Full Name: %s", age, fullName)

	return returnString
}

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	fmt.Println("This is foobar()")
}

// a function that accepts an int and returns a function
func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {

	myFirstFunction()
	functionWithArguments(5, 7)
	functionWithArguments2(1, 2, 3, 4.1, "akash")

	fmt.Println(f3(2.1))

	Sum, Product := f5(2, 3)
	fmt.Println(Sum, Product)

	fmt.Println(sum(1, 4))

	f1(1, 2, 3, 4, 5)
	f1() //with no arguments //output -> nil

	nums := []int{0, 1, 2, 3}

	f2(nums...)       // use elipsis to expand the variables
	fmt.Println(nums) // the nums value will be modified

	s, p := SumAndProduct(2.0, 5.1, 10.)
	fmt.Println(s, p)

	info := personInfo(30, "wolf", "asd", "kgf")
	fmt.Println(info)

	// defer statement execute in the reverse order
	// so first foobar will execute and then foo will execute

	// behind the defer statement are pushed on top of the stack
	// and pop one by one
	defer foo()
	defer foobar()
	bar()

	fmt.Println("This is just to test.")

	// anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("This is an anonymous function!")

	a := increment(10)
	fmt.Printf("%T #%v\n", a, a)
	a()
	fmt.Println(a())
}

// defer statement postpones the execution of a statement
// to the completion of surrounding function either through panik
// or completion

// defer is mainly used to clean up
// defer vs return
/*

We can break up defer ... into three parts.

invoke defer, evaluating the value of function parameter.
execute defer, pushing a function in to stack.
execute functions in stack after return or panic.

func test4() (x int) {
    defer func(n int) {
        fmt.Printf("in defer x as parameter: x = %d\n", n)
        fmt.Printf("in defer x after return: x = %d\n", x)
    }(x)

    x = 7
    return 9
}

invoke defer, evaluating the value of n, n = x = 0, so x as parameter is 0.
execute defer, pushing func(n int)(0) onto stack.
execute func(n int)(0) after return 9, n in fmt.Printf("in defer x as parameter: x = %d\n", n) has been evaluated, x in fmt.Printf("in defer x after return: x = %d\n", x) will be evaluated now, x is 9

in defer x as parameter: x = 0
in defer x after return: x = 9
in main: x = 9
*/
