package main

import "fmt"

func main() {
	p := fmt.Printf
	pl := fmt.Println

	title1, author1, year1 := "The devine comedy", "Dante", 1320
	title2, author2, year2 := "Macbeth", "William Shakespear", 1606

	type book struct {
		title, author string
		// author string
		year int
	}
	// order matters that is why not ideal technique
	mybook := book{title1, author1, year1} // struct literal
	pl(mybook)
	pl(mybook.title) // access field in structs

	// now order does not matter
	bestbook := book{title: title2, author: author2, year: year2}
	pl(bestbook)

	abook := book{year: 2021} // rest of the fields will take deafult value
	pl(abook)

	// mybook.page // this is will give error

	p("%#v\n", abook)

	// comparison of struts
	pl(abook.year == bestbook.year) // false

	bestbook.year = 2021 // the value of structs can be changed
	pl(bestbook)

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "diana",
		lastName:  "Muller",
		age:       30,
	}

	p("%#v\n", diana)

	// anonymous

	type Anaonymous struct {
		string
		float64
		bool
	}

	b1 := Anaonymous{"akash", 10.2, true}
	pl(b1)

	//nested structs
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "Akash",
		salary: 4000,
		contactInfo: Contact{
			email:   "akash@gmail.com",
			address: "India",
			phone:   91,
		},
	}

	p("%#v\n", john)

}
