package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("Cities is equal to nil: ", cities == nil)
	fmt.Println(len(cities))

	// cities[0] = "London" // this is give error because the len of slice is 0

	numbers := []int{1, 2, 3, 4}
	fmt.Println(numbers)

	//make can be used to create a slice of a fixed len
	nums := make([]int, 2)
	fmt.Println(nums) //[0, 0]

	//** COMPARING SLICES **//
	// a Go slice can only be compared to nil

	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	append_numbers := []int{}

	append_numbers = append(append_numbers, 10, 10, 20)
	fmt.Println(append_numbers)

	// append slice to slice
	n := []int{100, 200}
	append_numbers = append(append_numbers, n...) // use ellipsis
	fmt.Println(append_numbers)

	// make will create empty slice of same len
	dest_slice := make([]int, len(append_numbers))
	// dest_slice := []int{} // not work because len of dest_slice is 0
	// if we create dest_slice with smaller number than len(append_numbers) only that many numbers will be copied
	np := copy(dest_slice, append_numbers)

	fmt.Println(dest_slice, np)

	// slice expression
	a1 := []int{3, 7, 3, 7, 9, 3, 1} // slice will also work with array but only return a slice
	// a[start: stop] // this will include start and exclude stop
	b1 := a1[1:5] // b1 will always be a slice,a1 can be array or slice
	fmt.Printf("%v %T\n", b1, b1)

	c1 := a1[:4] // same as python
	_ = c1

	s1 := []int{1, 2, 3, 4, 5, 6}
	s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
	fmt.Println(s1)          // -> [1 2 3 4 100]

	s1 = append(s1[:4], 200) // overwrites the last element
	fmt.Println(s1)          // -> [1 2 3 4 200]

	var nums1 []int
	// nums1 :=  []int{}
	fmt.Printf("%#v\n", nums1)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums1), cap(nums1))

	nums1 = append(nums1, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(nums1), cap(nums1))

}
