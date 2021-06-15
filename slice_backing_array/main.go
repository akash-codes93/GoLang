package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// slice
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	v2, v3 := v1[0:3], v1[1:4]

	v2[1] = 600
	v1[2] = 800

	fmt.Println(v1, v2, v3) // all the three slices will change because the backing array is same.

	// v1, v2 and v3 share the same backing array but the starting header is not same
	fmt.Printf("%p %p %p\n", &v1, &v2, &v3)

	// array
	p1 := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	p2, p3 := v1[0:3], v1[1:4]

	p2[1] = 600
	p1[2] = 800
	// p1 is the backing array of p2 and p3 so 800 will be reflected to both p2 and p4
	// but 600 will only be reflected to p2 and p3
	fmt.Println(p1, p2, p3)

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda

	// slices see the whole backing array

	n1 := []int{10, 20, 30, 40, 50}
	n2 := n1[:3]
	n2[0] = 66
	fmt.Println(n1) //[66 20 30 40 50] This means n1 and n2 were using the same backing array

	n11 := []int{10, 20, 30, 40, 50}
	n22 := append(n11, 100)
	n22[0] = 66
	fmt.Println(n11) // [10, 20, 30, 40, 50] n11 will not change because append creates a new backing array

	// The append() function creates a new backing array with a larger capacity
	// to avoid creating a new backing array when the next append() is called.

	// coding challenge

	slice := []int{3, 8, 1}

	for index, value := range slice {
		println(index, value)
	}

	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6 // run time error will come, you need to use append

	a := 10.0
	mySlice[0] = a

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

}
