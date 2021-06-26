package main

import "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "bycycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // can use p.price = 300
	p.productName = "bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 12
}

func main() {
	p := fmt.Printf
	pl := fmt.Println

	var b *int
	a := 5
	b = &a
	c := &b

	pl(a, b, *b, &b)
	*b = *b + 1
	pl(a, b, *b, &b)

	//pointer to pointer
	pl(**c) // 6

	var ptr *float64
	pl(ptr) // nil

	ptr1 := new(int) // new can also be used to declare a pointer
	ptr1 = &a

	pl(ptr1)
	p("%p\n", ptr1) // %p is used to print a pointer

	//* COMPARING POINTER *//
	// two pointer are only same of they point to the same variable :: does not matter if they have the same value
	// or if thye are nil
	d := &a
	pl(d == b)

	var x int = 8
	ptr2 := &x

	pl("The value of x before func call: ", x)
	ptr3 := change(ptr2)
	pl("prt3: ", ptr3) // address of a float variable
	pl("The value of x after func call: ", x)

	change(&x)

	gift := Product{
		price:       100,
		productName: "Toffee",
	}
	pl("before calling change product: ", gift)
	// changeProduct(gift)
	changeProductByPointer(&gift)
	pl("after calling change product: ", gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	pl("after calling change prices: ", prices) // 2.3.4

	priceMap := map[string]int{
		"a": 10000,
		"n": 100,
	}
	changeMap(priceMap)
	pl("after calling change map: ", priceMap) // map[a:10 b:12 n:100]

	//array behaves like int :: pass by value
}
