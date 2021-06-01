package main

import (
	"fmt"
	// "math"
)

func main() {
	// constants need to be initialized when declared.
	const days int = 7

	fmt.Println(days)

	// const errors are detected at compile time

	// x, y := 5, 0
	// fmt.Println(x / y) // this error will be detected at runtime

	// const a = 5
	// const b = 0
	// fmt.Println(a / b) // this will be detected at complile time. vscode will show this as error

	// multiple constants
	const (
		min1 = 5
		min2 = 10
	)

	// in group constant the below constants repeate the above one
	const (
		min1_ = 6
		min2_
	)

	fmt.Println(min1, min2)
	fmt.Println(min1_, min2_) // output 5 5

	// cannot change value of a const
	// cannot initate a const at runtime
	// const belong to compile time

	// const temp = math.Pow(2, 5) // this will not work

	/*
		cannot use varibale to assign a const
		t := 5
		const temp = t // this will not work
	*/

	// The below will work
	const work = len("akash") // this will work because len function is available at compile time because it is a built-in function
	fmt.Println(work)

	const ale float64 = 5.1 // typed const
	fmt.Println(ale)

	const ab = 6 // untyped const
	fmt.Println(ab)

	const c = min1 * min2 // const expression ..evaluated at compile time.
	fmt.Println(c)

	const ct = "hello " + "go!"
	const cd = min1 > min2

	fmt.Println(ct, cd)

	const t = 5
	const t2 = 5.1
	const t3 = t * t2 // this is allowed because t and t2 are untyped const

	fmt.Printf("%T\n", t3) // float64

	const (
		zi = (iota * 2) + 1 // default 0
		_                   // blank identifier
		z2
		z3
	)
	print(zi, z2, z3)
	//There are ONLY boolean constants, rune constants, integer constants, floating-point constants, complex constants, and string constants.
}
