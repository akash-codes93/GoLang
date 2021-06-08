package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Long method
	// var a int
	// var b float64
	// var c bool
	// var d string

	// short method
	var (
		a int
		b float64
		c bool
		d string
	)

	_, _, _, _ = a, b, c, d

	x, y, z := 20, 15.5, "Gopher!"
	_, _, _ = x, y, z

	fmt.Println(a, b, c, d, x, y, z)

	const daysWeek = 7
	const lightSpeed = 299792458
	const pi = 3.14159

	_, _, _ = daysWeek, lightSpeed, pi

	// untyped
	const (
		daysWeek1   = 7
		lightSpeed1 = 29792458
		pi1         = 3.14159
	)

	_, _, _ = daysWeek1, lightSpeed1, pi1

	const secPerDay = 24 * 60 * 60
	const daysYear = 365

	fmt.Printf("%d\n", secPerDay*daysYear)

	// fmt
	x1, y1, z1 := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d, y is %f, z is %s\n", x1, y1, z1)
	fmt.Printf("score is %#v\n", score)

	// 2.
	fmt.Printf("z is %q\n", z1)

	// 3.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x1, y1, z1, score)

	// 4.
	fmt.Printf("y type: %T, score type: %T\n", y1, score)

	// operators
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
}
