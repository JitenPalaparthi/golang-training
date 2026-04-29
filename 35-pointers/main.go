package main

import "fmt"

func main() {

	var num1 int = 9897934

	fmt.Println("address of num1:", &num1)

	var ptr1 *int = &num1

	fmt.Println("address of num1 thru ptr1:", ptr1)

	*ptr1 = 87683234 // dereference
	fmt.Println(*ptr1, num1)

	var ptr2 **int = &ptr1

	var ptr3 ***int = &ptr2

	***ptr3 = 9999999 // dereference

	fmt.Println("num1:", num1)

	var ptr5 *int

	fmt.Println(ptr5)

	//ptr2 := new([]int)

	// slice1 := make([]int, 100)    // This could be on stack
	// slice2 := make([]int, 100000) // This could be on heap

}

// Pointer holds the address
// Pointer can be nil , that means it does not hold any address
// every data is stored in an address

//
