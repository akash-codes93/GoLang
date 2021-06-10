package main

import "fmt"

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
}
