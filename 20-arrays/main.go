package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr1 [5]int // zero values [0 0 0 0 0]

	fmt.Println(arr1)

	arr1[0] = 10
	arr1[1] = 20
	arr1[2], arr1[3], arr1[4] = 30, 40, 50

	fmt.Println(arr1, len(arr1))

	// type of array contains its length

	fmt.Println("type of arr1:", reflect.TypeOf(arr1))

	arr2 := [4]int{10, 11, 12, 13}

	arr3 := [...]int{10, 20, 30, 54, 46, 76, 434, 5, 65, 34} // evaluated by the compiler

	fmt.Println("type of arr2:", reflect.TypeOf(arr2), "len:", len(arr2), arr2)
	fmt.Println("type of arr3:", reflect.TypeOf(arr3), "len:", len(arr3), arr3)

	sum := 0

	for _, v := range arr3 {
		sum = sum + v
	}
	println(sum)

	// find max and min elements from an array

	// max := arr3[0]
	// min := arr3[0]

	// for _, v := range arr3 {
	// 	if max < v {
	// 		max = v
	// 	}

	// 	if min > v {
	// 		min = v
	// 	}
	// }

	max, min := GetMaxAndMinFromArray(arr3)

	println("max:", max, "min:", min)

	max, min = GetMaxAndMinFromArray4(arr2)

	println("max:", max, "min:", min)

}

// cannot type caste arrays

func GetMaxAndMinFromArray(arr [10]int) (max int, min int) {
	if len(arr) >= 0 {
		max := arr[0]
		min := arr[0]

		for _, v := range arr {
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

func GetMaxAndMinFromArray4(arr [4]int) (max int, min int) {
	if len(arr) >= 0 {
		max := arr[0]
		min := arr[0]

		for _, v := range arr {
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

// Due to these reasons , we generallyu dont use arraus as input parameters for functions
