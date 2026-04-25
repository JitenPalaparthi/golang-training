package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []int // this is slice declaration
	// Ptr = nil, Len = 0 Cap = 0

	if slice1 == nil {
		println("nil slice")
		// instantiate
		slice1 = make([]int, 10, 20) // Ptr= 0x1012312 Len = 10 Cap = 20
	}
	fmt.Println(slice1)

	for i := range slice1 {
		slice1[i] = rand.IntN(1000)
	}

	fmt.Println(slice1)

	//var str1 string // Ptr:nil Len:0 ==""

	max, min := GetMaxAndMin(slice1)

	println("Slice1 max:", max, "min:", min)

	slice2 := []int{1, 23, 344, 4, 34, 6775, 234, 436, 544, 213, 4656, 9}

	max, min = GetMaxAndMin(slice2)

	println("max:", max, "min:", min)

	slice3 := make([]int, 5)
	slice3[0], slice3[1], slice3[2], slice3[3], slice3[4] = 43, 390, 34, 2, 4

	max, min = GetMaxAndMin(slice3)

	println("max:", max, "min:", min)

	slice4 := []int{} // Ptr: Zerobase pointer, Len: 0 Cap: 0 , slice is not nil
	max, min = GetMaxAndMin(slice4)

	println("max:", max, "min:", min)

}

func GetMaxAndMin(slice []int) (max int, min int) {
	if len(slice) > 0 {
		max = slice[0]
		min = slice[0]

		for _, v := range slice {
			if max < v {
				max = v
			}

			if min > v {
				min = v
			}
		}
	}
	return max, min
}

// what is called a slice
// slice is a growable array
// a slice can be nil
// a slice has a header
// Ptr unsafe.Pointer // 8 on 64bit machines
// Len int // 8 on 64bit machines
// Cap int // 8 on 64bit machines
// To instantiate a slice can use make function
