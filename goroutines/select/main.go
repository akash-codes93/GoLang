package main

import (
	"fmt"
	"time"
)

func main() {
	pl := fmt.Println

	start := time.Now().Local().UnixNano() / 1000000

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salut"
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 2; i++ {
		select { // just like switch but used with channels
		case msg1 := <-c1:
			pl("Revieved:", msg1)
		case msg2 := <-c2:
			pl("Revieved:", msg2)

		default:
			pl("No activity")
		}

	}

	end := time.Now().Local().UnixNano() / 1000000
	pl(end - start) // it takes close to 2 sec to execute as both the sleeps eecuted concurrently

}
