package main

import (
	"fmt"
	"unsafe"
)

// raw string = `this is a raw string`

func main() {
	text := "यह भारत 1 number है"
	// unicode has 144k characters which alone or combine to make a graphene
	// unicode has different encoding scheme to store it has binary like utf-8, utf-32 etc
	// ascii and unicode chars matches to have compatibility
	// भा : this is a graphene which is made of two unicode characters भ and ा and it takes two runes to represent
	// rune == unicode (incase if utf-8) and graphene may take multiple rune to represent

	fmt.Println("Length of text: ", len(text)) // -> 35 len return total number of bytes in the string
	// iterating over this text means iterating over unicode/rune
	// index shows index of each rune which is utf-8 encode which is variable length encoded
	for idx, e := range text {
		fmt.Printf("idx: %d, graphene: %c, type: %T, size: %d, value: %v\n", idx, e, e, unsafe.Sizeof(e), e) //value: -> 2351 unicode code point
	}

	// lets take a look at index of this.
	fmt.Println("Indexing 2:8 -> ", text[2:8]) // -> �ह � output because it is byte indexing

	// ***** fundamental -> when you loop you loop over rune(unicode) which is of variable length
	// but indexes are at byte level text[1] : at 1st byte value

	// each char is a rune

	fmt.Println("-------Rune-------")
	// rune is alias for int32 which is 4 byte
	// some time a graphene may take two runes to store eg. है
	r := []rune(text)
	fmt.Println(r, len(r)) // list of unicodes[rune] and total runes

	// byte conversion
	fmt.Println("-------bytes-------")
	b := []byte(text)
	// this is printing byte by byte
	for idx, e := range b {
		fmt.Printf("idx: %d, graphene: %c, type: %T, size: %d\n", idx, e, e, unsafe.Sizeof(e))
	}
	fmt.Println("Overall:", b)

}
