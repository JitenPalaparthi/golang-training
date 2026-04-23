package main

import "fmt"

func main() {

	arr2d := [2][2]int{{1, 2}, {3, 4}}

	for _, a1 := range arr2d {
		for _, v := range a1 {
			print(v, " ")
		}
		println()
	}

	arr3d := [2][2][3]int{{{1, 2, 3}, {4, 5, 6}}, {{7, 8, 9}, {10, 11, 12}}}

	sum := 0

	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				sum += v
			}
		}
	}

	println("sum:", sum)

	// any array

	arr1 := [10]any{10, 10.1, uint8(20), uint16(343), true, "hello Wrold", any(100), any(1212.123)}
	fmt.Println(arr1)
}

// clears the array
func clear(arr [5]int) {
	for i := range arr {
		arr[i] = 0
	}
}
