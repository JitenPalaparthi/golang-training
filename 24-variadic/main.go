package main

import "fmt"

func main() {

	fmt.Println(1, 2, "true", "hello world", "Jiten", "Sai")

	slice1 := []int{}
	slice1 = append(slice1, 10, 234, 4352, 234, 324, 4, 4, 433445, 45, 3453, 345, 4534, 645, 324, 4535)

	sum := SumOf()
	fmt.Println("Sum:", sum)

	sum = SumOf(343)
	fmt.Println("Sum:", sum)

	sum = SumOf(343, 234, 23)
	fmt.Println("Sum:", sum)

	sum = SumOf(10, 234, 4352, 234, 324, 4, 4, 433445, 45, 3453, 345, 4534, 645, 324, 4535)
	fmt.Println("Sum:", sum)

	sum = MulAndSumOf(2, 10, 20)
	fmt.Println("Sum:", sum)

	sum = SumOf(slice1...) // a slice can be converted as a variadic argument
	fmt.Println("Sum:", sum)
	arr1 := [5]int{10, 11, 12, 13, 14}

	slice2 := arr1[:] // an array can be converted into a slice

	sum = SumOf(slice2...)
	fmt.Println("Sum:", sum)

	sum = SumOf(arr1[:]...)
	fmt.Println("Sum:", sum)

}

// variadic functions
// variadic paramater must be the last parameter to the func
// variadic patameter cannot be used other than function or methods

func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

//func MulAndSumOf( nums ...int,mul int) int { // variadic parameter must be the last parameter

func MulAndSumOf(mul int, nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v * mul
	}
	return sum
}
