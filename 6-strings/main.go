package main

import (
	"fmt"
	"math/rand/v2"
	"unsafe"
)

func main() {

	str1 := "Hello 世界 ❤️!" // It is not 11 bytes(may be 11 chars) , shorthand notation

	str1 = "世界 ❤️!" // it creates a new memory and gives that address to the pointer in string structure

	var str2 string = "Hello World!" // general way of creating a variable and assigning a value

	var str3 = "Hello Universe!" // type inference

	var str4 string // just declaring a string

	str4 = "Hello Golang Folks!" // mutating or assigning a value to the string

	// What is the structure of the string
	// string is collection of bytes in go
	// Size and Len of the string
	// how strings work behind the scene
	// pointer
	// nil

	// i := 100

	// j := 5466.435

	// k := true

	r := rand.IntN(10000000000) // runtime

	println(r)

	println(str1, "len of str1:", len(str1), "Size of the str1:", unsafe.Sizeof(str1))
	println(str2, "len of str2:", len(str2), "Size of the str2:", unsafe.Sizeof(str2))
	println(str3, "len of str3:", len(str3), "Size of the str3:", unsafe.Sizeof(str3))
	println(str4, "len of str4:", len(str4), "Size of the str4:", unsafe.Sizeof(str4))

	// fmt.Println("address of str1:", byte(str1[0]))

	str1 = fmt.Sprint(r) // It converts the int to string and assign the value to the existing string
	// the above is not known to the compiler only based on runtime it knows
	println(str1, "len of str1:", len(str1), "Size of the str1:", unsafe.Sizeof(str1))

	var str5 string // default or zero value
	println(str5, "len of str5:", len(str5), "Size of the str5:", unsafe.Sizeof(str5))

}

// Study UTF-8 Chars
// Study about pointers

// string structure, string header
// Pointer --> The address to the original data, 8/4 bytes (depends on the type of the machine arch 64bit or 32bit)
// Length  --> int , 8 bytes or 4 bytes

// string literals are stored in ro data segment
// where ever there is a pointer, it comes with nil --> 0x00 ,0
