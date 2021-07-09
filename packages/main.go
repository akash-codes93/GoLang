package main

import (
	"fmt"
	"mypackage/greet"
	"mypackage/numbers" // package is to be added in GOROOT folder // use go env to find out path
)

// any executable package is called main

func main() {

	greet.PleaseGreet()

	pl := fmt.Println
	n := 101

	status := numbers.Even(n)
	pl(status)

	status = numbers.IsPrime(int(n))

	pl(status)

	pl(numbers.OddAndPrime(25))
	pl(numbers.OddAndPrime(13))

}

/*********  Notes from packages file ***********/
// all files in a directory must belong to the same package

// function name should start with capital letter

// lower case function becomes private to the package :: cannot be called from outside
