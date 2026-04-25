package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice2 := slice1
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	slice3 := slice1[:] // all the elements
	fmt.Println("slice3:", slice3, "len:", len(slice3), "cap:", cap(slice3))

	slice4 := slice1[:5] // but not 5
	fmt.Println("slice4:", slice4, "len:", len(slice4), "cap:", cap(slice4))

	slice5 := slice1[3:8] // but not 8
	fmt.Println("slice5:", slice5, "len:", len(slice5), "cap:", cap(slice5))

	slice6 := slice1[5:]
	fmt.Println("slice6:", slice6, "len:", len(slice6), "cap:", cap(slice6))

	slice5 = append(slice5, 99999)
	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("slice5:", slice5, "len:", len(slice5), "cap:", cap(slice5))
	fmt.Println("slice6:", slice6, "len:", len(slice6), "cap:", cap(slice6))

	// remove an element from the slice is always adjusting the slice with index
	slice7 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice8 := append(slice7[:4], slice7[5:]...)
	fmt.Println("slice8:", slice8, "len:", len(slice8), "cap:", cap(slice8))
}
