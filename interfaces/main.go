package main

import (
	"fmt"
	"math"
)

/*
Imp: so if a type implements these two methods
it automatically implements the interface
You don't have to explicitly mention it in the type
*/
type shape interface {
	area() float64
	perimeter() float64
}

/****** The approaches is not good ******/
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// a type can have also other methods that don't belong to the interface.
func (c circle) volume() float64 {
	return (4 / 3) * (math.Pi * math.Pow(c.radius, 3))
}

func printCircle(c circle) {
	fmt.Println("Shape: ", c)
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perimeter())
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func printRectangle(r rectangle) {
	fmt.Println("Shape: ", r)
	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.perimeter())
}

/****** The better approach :: Interfaces  ******/

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %#v\n", s.area())
	fmt.Printf("Perimeter: %#v\n", s.perimeter())
}

type emptyInterface interface {
}

type person struct {
	info interface{}
}

/*
interfaces only contains the signature of the methods and not the implementations
so rectangle and circle type automatically implements the shape interface.
because they both have the area and perimeter methods
Also, a type must implement all the interfaces methods, partial will also not work
*/
func main() {
	p := fmt.Printf
	pl := fmt.Println

	// p("%T\n")

	c1 := circle{radius: 5.}
	r1 := rectangle{width: 3., height: 2.1}

	printCircle(c1)
	printRectangle(r1)
	/***** same as the above *****/
	print(c1)
	print(r1)

	var s shape = c1
	pl(s)

	var newShape shape = circle{radius: 2.}
	p("%T\n", newShape)

	// newShape.volume() // this will be error because we can only use methods defined inside  interface

	c2 := circle{radius: 2.1}
	pl(c2.volume()) // this is possible

	// type assertion
	value, success := newShape.(circle) // the type of newshape is changed to circle

	if success {
		p("Volume of ball: %v\n", value.volume()) // now this will not give error
	}

	/*** embedded interface
	An interface contains other interface
	circular embedded is not allowed
	***/
	var empty interface{}

	empty = 5
	pl(empty)

	empty = "GO"
	pl(empty)

	empty = []int{4, 5, 6}
	pl(empty)
	// pl(len(empty)) // This will give error :: assertion required
	fmt.Println(len(empty.([]int))) // This will work

	you := person{}
	// now info can have any value
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5.6, 7.8}
	pl(you.info)
}
