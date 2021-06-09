package main

import (
	"fmt"
	"time"
)

func main() {

	input := 3

	switch input {
	case 1:
		fmt.Println("This is 1")
		// break // break is default in Go not need to put it explicitly
	case 10, 5:
		fmt.Println("This is 10")

	default:
		fmt.Println("This is default statement")
	}

	switch true {
	case input%3 == 0:
		fmt.Println("This is divisibe by 3")
	}

	hour := time.Now().Hour()
	_ = hour
}
