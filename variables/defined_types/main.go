package main

import (
	"fmt"
)

type km float64
type mile float64

func main() {

	type speed uint

	var s1 speed = 120
	var s2 speed = 100

	fmt.Println(s1 - s2)

	var x uint
	// x = s1 // this will give complie time error as x and s1 are not of same type
	x = uint(s1)
	_ = x

	var s3 speed = speed(x) // both side type case is available
	_ = s3

	// type casting can only be done if they share same underlying type

	var parisToLondon km = 4513
	var parisToLondonMiles mile

	parisToLondonMiles = mile(parisToLondon / 0.621)

	fmt.Println(parisToLondonMiles)

	// alias
	//  byte and uint8 is same
	// prune and int32 are same

	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	//no need to convert operations (same type)
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60
}
