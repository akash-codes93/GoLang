package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// time.Sleep(time.Second * 2)
	}
	fmt.Println("f1 execution stopped")
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

	go f1()
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine()) // 2

	f2()

	time.Sleep(time.Second * 2) // if this sleep is removed f1 ouptut will not be seen
	// this means that while the main code sleep the go routine f1 can be exected concurrently
	// if we put sleep at f1 loop iteration this means that only all iteration are complete before main sleep if revoked.
	// meaning that main f1 routine only executes till the time main rountine waits
	fmt.Println("main execution stopped")
}
