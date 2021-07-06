package main

import (
	"fmt"
	"time"
)

func main() {
	pl := fmt.Println
	p := fmt.Printf

	c2 := make(chan int, 3) // 3 is the channel capacity // buffered channel
	_ = c2

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			p("func goroutine #%d starts sending data into the channel\n", i)
			c <- 10
			p("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c) // close channel when you want to tell the reciever that all data is send.
	}(c2)

	pl("main goroutine sleeps for 2 seconds")
	time.Sleep(2 * time.Second)

	for v := range c2 { // v := <- c
		pl("Value revieved from channel: ", v)
	}

	// a panic is created when you try to send on a close channel

}
