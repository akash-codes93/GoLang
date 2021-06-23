package main

import "fmt"

func main() {
	p := fmt.Printf
	pl := fmt.Println

	var employees map[string]string
	pl(employees)
	p("%#v\n", employees)
	pl(employees["Dan"]) // will not give error - print ""

	// map keys can only be comparable type

	// employees["Dan"] = "Akash" // you cannot assign value to nil map

	people := map[string]float64{}
	people["John"] = 64
	people["Marry"] = 10

	pl(people)

	map1 := make(map[int]int) // initialize an empty map using map
	map1[4] = 8
	pl(map1)

	balances := map[string]float64{
		"USD": 123.12,
		"EUR": 2736, // comma is required when you use multiple lines
	}

	pl(balances)
	pl(balances["RON"]) // 0  as default value for float64 is 0

	v, ok := balances["RON"] // ok signifies that RON key exists in the map
	pl(v, ok)

	// looping over map
	for k, v := range balances{
		pl(k, v)
	}

	delete(balances, "USD") // will delete the key from the map. No need to check if key exists in the map
	pl(balances)

	// two maps cannot be compared, a map can only be compared to nil

	// string representation of a map
	// can only be used only when both key and value both are string

	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbors := friends

	friends["Dan"] = 50
	pl(neighbors) // the neighbour map will also change because both the map share the same map headers

	newFriends := make(map[string]int)

	for k, v := range friends{
		newFriends[k] = v
	}

	pl(newFriends)
	 
	// invalid map as there is duplicate key, compile time error
	// balances := map[string]float64{
    //     "A": 201.4,
    //     "B": 201.4,
    //     "B": 600.4,
    // }
}
