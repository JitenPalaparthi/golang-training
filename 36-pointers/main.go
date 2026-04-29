package main

import (
	"fmt"
	"unsafe"
)

func main() {

	slice1 := []int{10, 20} // len 2 cap 2

	Append(slice1, 30, 40, 50, 60)

	fmt.Println("no change to the original slice", slice1)

	slice1 = AppendR(slice1, 30, 40, 50, 60) // return a slice is nothing but the heards of inside func are copied to the return variable

	fmt.Println("changes the original slice buy using return", slice1)

	slice2 := []int{10, 20} // len 2 cap 2
	AppendP(&slice2, 30, 40, 50, 60)

	fmt.Println("changes the original slice using pointers", slice2)

	var ok bool = true
	var ptrok *bool = &ok

	fmt.Println("ok size:", unsafe.Sizeof(ok))       //1 byte
	fmt.Println("ptrok size:", unsafe.Sizeof(ptrok)) // 8 bytes

	arr1 := [5]int{10, 20, 30, 40, 50}

	ptrArr1 := &arr1                                     // *[5]int
	fmt.Println("ptrArr1 size:", unsafe.Sizeof(ptrArr1)) // 8 bytes

}

func Append(slice []int, nums ...int) {

	for _, v := range nums {
		slice = append(slice, v) // append chagnes the slice header
	}

}

func AppendR(slice []int, nums ...int) []int { // call by values
	for _, v := range nums {
		slice = append(slice, v) // append chagnes the slice header
	}
	return slice
}

func AppendP(slice *[]int, nums ...int) { // call or pass by reference
	fmt.Println("Size of ptr slice:", unsafe.Sizeof(slice))
	for _, v := range nums {
		*slice = append(*slice, v) // append chagnes the slice header
	}
}
