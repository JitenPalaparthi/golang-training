package main

import "errors"

func main() {

	// num, div := 100, 0

	// println(num / div) // very common error, panic which we call it as divide by zero

	// arr := [2]int{10, 20}

	// for i := 0; i <= len(arr); i++ {
	// 	println(arr[i]) //index out of range [2] with length 2
	// }

	var ptr *int // it is a pointer, which should an address, but it is nil

	println(*ptr) // dereferencing the pointer,  invalid memory address or nil pointer dereference

	num := 100

	err := incr(&num)
	if err != nil {
		println(err.Error())
	} else {
		println(num)
	}

	var ptr1 *int

	err = incr(ptr1)

	if err != nil {
		println(err.Error())
	} else {
		println(*ptr)
	}

	println("Hello Main")

}

// errors  --> values
// panic panics the execution
// panic crashes the application by default

// to handle the nil pointer dereferecing through error handingling and avoid panic
func incr(ptr *int) error {
	if ptr != nil {
		*ptr++
		return nil
	}
	return errors.New("invalid pointer")
}
