package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	var slice1 []int // nil slice

	// fill the slice with rand values

	for range 100 {
		slice1 = append(slice1, rand.IntN(100)) // append can also be used on a nil slice
	}

	fmt.Println(slice1)

	// find the duplicate values along with their count

	map1 := make(map[int]int)

	for _, v := range slice1 {
		// c := map1[v] // either v is there or not there
		// map1[v] = c + 1

		c, ok := map1[v] // either v is there or not there
		if !ok {
			map1[v] = 1
		} else {
			map1[v] = c + 1
		}
	}

	fmt.Println(map1)

}
