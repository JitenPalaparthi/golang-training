package main

import "fmt"

func main() {

	num := 0b1100110011001100 // binary format
	mask := 0b1111

	result := byte(num & mask)

	// 1100110011001100
	// 						&
	// 0000000000001111
	//------------------
	// 0000000000001100
	fmt.Printf("%b\n", result)

	// To On a bit

	position := 8
	r := num | (1 << position) // 1000000000
	fmt.Printf("%b\n", r)

	//1100110011001100
	//0000001000000000
	//1100111011001100

}
