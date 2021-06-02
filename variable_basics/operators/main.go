package main

import (
	"fmt"
	"strconv"
)

func main() {
	x, y := 5, 6

	x += y
	fmt.Println(x, y)

	counter := 0
	// counter++ this is an statement and not an operator
	counter++ // pre-operators does not exists ++ counter -- error
	fmt.Printf("%d\n", counter)

	a, b := 5, 10
	fmt.Println(a == b)

	// && ||
	// in case of && if first expression is false, then second expression will not be evaluated
	fmt.Println(!(a > 0))

	s := string(99)
	fmt.Println(s) // output c as 99 corresponding to string is c

	// using fmt
	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)
	fmt.Printf("%T\n", myStr) // string

	var myStr1 = fmt.Sprintf("%d", 23456)
	fmt.Println(myStr1)
	fmt.Println(string(23456)) // this will output unicode character

	s1 := "3.1234"
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)

}
