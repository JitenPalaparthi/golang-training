package main

import "fmt"

func main() {

	slice1 := []int{10, 11, 12, 13, 14}

	println("len of slice1:", len(slice1), "Capacity of slice1:", cap(slice1))

	slice1 = append(slice1, 15) // the capacity is doubled
	println("len of slice1:", len(slice1), "Capacity of slice1:", cap(slice1))

	slice1 = append(slice1, 16, 17, 18, 19, 20)
	println("len of slice1:", len(slice1), "Capacity of slice1:", cap(slice1))

	slice2 := make([]int, 11, 20) // can you always give the cap no..

	for i := 0; i <= 10; i++ {
		slice2[i] = i + 10
	}

	fmt.Println(slice2)
	println("len of slice2:", len(slice2), "Capacity of slice2:", cap(slice2))

}
