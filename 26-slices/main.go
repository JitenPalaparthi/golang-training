package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // len 10 cap 10
	fmt.Println("Ptr of slice1:", &slice1[0], "Len:", len(slice1), "Cap:", cap(slice1))
	fmt.Println(slice1)
	sum := SumOf(slice1)
	println("Sum:", sum)

	sum = SumOfAfterSq(slice1)
	println("Sum:", sum)

	fmt.Println(slice1)

	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // len 10 cap 10
	sum = SumOf(slice2)
	println("Sum:", sum)
	fmt.Println(slice2)

	sum = SumOfSqAddMore(slice2, 11, 12, 13, 14, 15)
	println("Sum:", sum)
	fmt.Println(slice2)

	// n := 10

	// s := Sq(n)

	// println("Sq:", s, "n:", n)

}

func Sq(num int) int {
	num = num * num
	return num
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumOfAfterSq(slice []int) int {
	sum := 0
	for i, v := range slice {
		slice[i] = v * v
		sum += slice[i]
	}
	return sum
}

func SumOfSqAddMore(slice []int, nums ...int) int {
	slice = append(slice, nums...) // The slice header is changed
	sum := 0
	for i, v := range slice {
		slice[i] = v * v
		sum += slice[i]
	}
	return sum
}

// append would change the header for 100% ,but what does it change , depends on the cap of the slice
// with in the capacity, it changes only the length
// beyond the capacity , it doubles(latest algo is based on the size 2X or 1/4X) and reassign the slice,
// so all three of them ptr, len, cap are changed
