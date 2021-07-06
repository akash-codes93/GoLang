package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	p := fmt.Printf
	pl := fmt.Println

	c := make(chan int)

	// only for receiving
	c1 := make(<-chan string)

	// only for sending
	c2 := make(chan<- string)

	p("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)

	n := <-c

	pl("value receivied: ", n)
	pl("Exiting main.....")

	close(c)
}
