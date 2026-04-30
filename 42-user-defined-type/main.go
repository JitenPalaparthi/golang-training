package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var e1 E1
	var e2 E2

	fmt.Println("Size of E1 variable:", unsafe.Sizeof(e1))
	fmt.Println("Size of E2 variable:", unsafe.Sizeof(e2))

	ep1 := new(E1) // Size of E1 varaible is 0 , no allocation
	ep2 := new(E2) // Size of E2 variable is 0 . no allocation

	if ep1 == nil {
		println("ep1 is nil")
	}

	if ep2 == nil {
		println("ep2 is nil")
	}

	fmt.Printf("address of ep1:%p\n", ep1) // Zerobase address
	fmt.Printf("address of ep2:%p\n", ep2)

	// num := 100

	// ptr := &num // pointer holds the address

	var empty struct{} // variable of empty struct

	empty = struct{}{}

	fmt.Println(empty)
	fmt.Printf("address of empty:%p\n", &empty)

}

// Empty structs

type E1 struct{}

type E2 struct{}

// What is zerobase address? If a type is empty , size if 0 , yet Go does not make the pointer nil , instead it uses a default address called zerobase address, or a dummy address
