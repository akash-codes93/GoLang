package main

import "fmt"

func main() {

	// p := fmt.Printf
	pl := fmt.Println

	var ch chan int

	pl(ch)

	ch = make(chan int)
	pl(ch)

	c := make(chan int)
	pl(c)

	// <- channel operator

	// SEND
	c <- 10

	// RECIEVE
	num := <-c
	pl(<-c)

	_ = num
	close(c)
}
