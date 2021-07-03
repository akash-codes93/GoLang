package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// data race

	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	//1.
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			//2.
			m.Lock()
			n++
			m.Unlock() // the code between lock and unlock will execute by one one goroutine at a time.
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock() // same as the above
			n--
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value of n: ", n) // the value of n will not be fixed
	/****  race detector  ****/
	// go run -race main.go

}
