package main

import (
	"fmt"
	"os"
)

func main() {
	/* os.Args is a slice of strings with 1 parameter
	as name of compiled file and
	rest are variables passed from cli*/
	fmt.Println("Arguments", os.Args)
	fmt.Println(len(os.Args))
}
