package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	slice1 := make([]int, 10)
	FillSlice(slice1)

	fmt.Println(slice1)
	fmt.Println("Ptr of slice1:", &slice1[0], "Len:", len(slice1), "Cap:", cap(slice1)) // 0x559a2152a000

	slice1 = append(slice1, 60)
	fmt.Println("Ptr of slice1:", &slice1[0], "Len:", len(slice1), "Cap:", cap(slice1)) // 0x559a2152e000

	slice1 = append(slice1, 70)
	fmt.Println("Ptr of slice1:", &slice1[0], "Len:", len(slice1), "Cap:", cap(slice1)) // 0x559a2152e000

	slice1 = append(slice1, 80, 90, 100, 110, 120, 130, 140, 150, 160)
	fmt.Println("Ptr of slice1:", &slice1[0], "Len:", len(slice1), "Cap:", cap(slice1)) // 0x559a2152e000

	fmt.Println(slice1)

}

// Slice has a header
// ptr len cap

func FillSlice(slice []int) {
	for i := range slice {
		slice[i] = rand.IntN(9999)
	}
}
