package main

import (
	"fmt"
	"time"
)

//name type

type names []string

// methods associated to type names :: reviever method
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func main() {
	p := fmt.Printf
	pl := fmt.Println

	const day = 24 * time.Hour

	p("%T\n", day)

	seconds := day.Seconds() // seconds is a reciever function
	p("%T\n", seconds)
	pl(seconds)

	friends := names{"ant", "marry"}
	friends.print()

	names.print(friends) // same as above :: not very common
	/*
		methods is a function that takes reciever as argument
		methods belongs to a type whereas a function belongs to a package
	*/

	mycar := car{
		brand: "Audi",
		price: 40000,
	}

	changeCar(mycar, "renault", 20000)
	pl("mycar after calling changecar: ", mycar) // no change :: {Audi 40000} :: pass by value

	mycar.changeCar1("oppel", 21000)
	pl("mycar after calling changecar1: ", mycar) // no change :: {Audi 40000} :: pass by value

	mycar.changeCar2("suzuki", 35000)
	/** same as above **/
	(&mycar).changeCar2("bmw", 450000)
	pl("mycar after calling changecar2: ", mycar) // value change :: {suzuki 35000} :: pass by value

	/**** more ways ****/
	var yourCar *car
	yourCar = &mycar
	fmt.Println(*yourCar)

	yourCar.changeCar2("Hyundai", 32000)
	pl(*yourCar)

	// we can create methods only for value types and not for pointer types

}
