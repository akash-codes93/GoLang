package main

import (
	"fmt"
	"time"
)

func main() {
	pl := fmt.Println

	c1 := make(chan int) //unbuffered channel

	c2 := make(chan int, 3) // 3 is the channel capacity // buffered channel
	_ = c2

	go func(c chan int) {
		pl("func goroutine starts sending data into the channel")
		c <- 10
		pl("func goroutine after sending data into the channel")
	}(c1)

	pl("sleep: ")
	time.Sleep(time.Second * 2)
	pl("revieved")
	d := <-c1
	pl(d)

	time.Sleep(time.Second * 2)

}
