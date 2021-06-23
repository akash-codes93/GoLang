package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// default 'utf-8' encoded
	// use single quote for strings

	fmt.Println("He says: \"Hello\"")
	fmt.Println(`He says: "Hello"`) // can use backticks to escape ""

	fmt.Println("akash,\nakashnew")

	fmt.Println(`
akash,
akashnew	
	`)

	var s1 = "akashnew"
	fmt.Println(s1[0]) // 97 because string is a slice of ascii codes

	// s1[0] = "x" // this will give error as string as immutable in python

	var1, var2 := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d\n", var1, var2) // Type: int32, Value: 98

	str := "tara"
	fmt.Println(len(str))

	fmt.Println(str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println()

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println()

	for _, r := range str {
		fmt.Printf("%c", r)
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(str))

	//slicing of strings
	// String slicing returns bytes, not runes. s1[0:3] means bytes at index 0, 1 and 2. Non-ascii letters are represented on more than one byte.

	s3 := "abcdefghijkl"
	fmt.Println(s3[2:5]) // -> cde, bytes from 2 (included) to 5 (excluded)

	s4 := "中文维基是世界上"
	fmt.Println(s4[0:2]) // -> � - the unicode representation of bytes from index 0 and 1.

	// returning a slice of runes
	// 1st step: converting string to rune slice
	rs := []rune(s4)
	fmt.Printf("%T\n", rs) // => []int32

	// 2st step: slicing the rune slice
	fmt.Println(string(rs[0:3])) // => 中文维

	// coding exercise

	var name = "akash"
	country := "India"
	p := fmt.Printf
	pl := fmt.Println
	p("Your name: %s\nCountry: %s\n", name, country)

	r := []rune("ă")
	pl(r)

	ma := "ma"
	m := "m"

	mother := ma + m + string(r)
	pl(mother)

	q3 := "țară means country in Romanian"

	for _, r := range q3{
		p("%c", r)
	}
	pl()

	s := "你好 Go!"

	// converting string to byte slice
	b := []byte(s)

	// printing out the byte slice
	fmt.Printf("%#v\n", b)

	// iterating over the byte slice and printing out each index and byte in the byte slice
	for i, v := range b {
		fmt.Printf("%d -> %d\n", i, v)
	}

}
