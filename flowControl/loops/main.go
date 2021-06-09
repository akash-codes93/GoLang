package main

import "fmt"

func main() {
	// i is block variable exists only inside the for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	// while loop : Note go does not have while keyword, you can use for loop to create while effect
	j := 10
	for j > 0 {
		fmt.Println(j)
		j--
	}
	 // handling of multiple variables in a for loop
	 for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
        fmt.Printf("i = %v, j = %v\n", i, j)
    }
}
