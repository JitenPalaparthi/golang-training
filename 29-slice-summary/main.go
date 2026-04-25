package main

import "fmt"

func main() {

	var slice1 []int // slice is nil
	if slice1 == nil {
		println("slice1 is nil")
	}

	// instantiate a slice using make function

	slice1 = make([]int, 10) // instantiate a slice

	fmt.Println(slice1) // zero values

	slice2 := []int{} // slice is not nil , len 0 cap 0 yet ptr: a dummy pointer zerobase pointer

	if slice2 == nil {
		println("slice2 is nil")
	} else {
		println("this concept slice2 is not nil, thought len is", len(slice2))
	}

	var slice3 []string // this is nil, even a nil slice can be appended

	slice3 = append(slice3, "Hello World", "How are you doing")

	fmt.Println(slice3)

	// make
	// append
	// copy
	// clear

}
