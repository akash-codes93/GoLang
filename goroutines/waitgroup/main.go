package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) { // wg is passed as a pointer
	fmt.Println("f1 execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second * 2)
	}
	fmt.Println("f1 execution stopped")
	//3.
	// Before exiting, call wg.Done() in each goroutine
	// to indicate to the WaitGroup that the goroutine has finished executing.
	wg.Done()
}

func f2() {
	fmt.Println("f2 execution started")
	for i := 5; i < 9; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution stopped")
}

func main() {
	// p := fmt.Printf
	pl := fmt.Println

	pl("Main execution started")
	pl("No. of cpus: ", runtime.NumCPU())
	pl("No. of Goroutines: ", runtime.NumGoroutine())

	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("OS: ", runtime.GOARCH)

	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))

	// 1.
	// Create a new instance of sync.WaitGroup (we’ll call it symply wg)
	// This WaitGroup is used to wait for all the goroutines that have been launched to finish.
	var wg sync.WaitGroup
	// 2.
	// Call wg.Add(n) method before attempting to
	// launch the go routine.
	wg.Add(1)

	go f1(&wg)
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine()) // 2

	f2()

	// 4.
	// Finally, we call wg.Wait()to block the execution of main() until the goroutines
	// in the WaitGroup have successfully completed.

	wg.Wait()
	fmt.Println("main execution stopped")
}
