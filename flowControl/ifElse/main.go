package main

import "fmt"

func main() {

	price, inStock := 80, true

	if price > 60 {
		fmt.Println("Too expensive")
	}
	// _ = inStock

	if price < 100 && inStock {
		fmt.Println("Available!")
	}

// 	if price { // valid in python but not in go
// 		fmt.Println("test")
// 	}

	if price < 100 {
		fmt.Println("it's cheap")
	} else if price == 80 {
		fmt.Println("on the edge")
	} else{
		fmt.Println("no")
	}


}

