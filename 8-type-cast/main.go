package main

import "fmt"

func main() {

	// 10 --> 10 // 0-9
	// 10 --> 0b1010 // 0 and 1
	// 10 -> A //  0-9 A B C D E F

	var num1 int = 10

	fmt.Printf("num1:%d num1:%b num1:%X\n", num1, num1, num1)

	var num2 uint8 = uint8(num1)

	print(num2)

	// var ok1 bool = true

	//var num3 uint8 = uint8(ok1) // a bool cannot be directly type casted to any other type

	// var num3 uint8

	// if ok1 {
	// 	num3 = 1
	// } else {
	// 	num3 = 0
	// }

	var num3 uint = 23432423 // 8 bytes 10101100101 10001100 11100111

	fmt.Printf("%b\n", num3)

	var num4 uint8 = uint8(num3) // (uint8)num3 1 byte 11100111 --> 11100111

	var num5 uint16 = uint16(num3) // 10001100 11100111. --> 10001100 11100111

	println(num4)
	fmt.Printf("%b\n", num4)

	println(num5)
	fmt.Printf("%b\n", num5)

	var float1 float32 = 13112.12 // 4 bytes 1 bit is sign 8 bits are exponent ,23 bits mantissa
	fmt.Printf("%b\n", float1)

	var num6 uint32 = uint32(float1) // lose the .12

	println(num6)

	var num7 int = 123123213

	var float2 float64 = float64(num7)

	fmt.Printf("%.2f", float2)

}

// numbering systems --> binary, decimal , hexa
// can type cast converts every type ? No
// among numbers yes --> int,uint, uint8-uint64, int8-int64, float32,float64
// there is no implicit type casting in go
