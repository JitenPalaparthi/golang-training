package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type char = int32 // now char is a type, not a new type just an another name to int32

func main() {

	var char1 rune = 'A' // rune --> int32

	var char2 int32 = 'B' // int32

	var char3 int64 = 'C' // int64

	var char4 uint64 = '世' // int64

	var char5 char = 'X'

	var char6 byte = 'A' // '世' // uint8

	fmt.Println(char1, char2, char3, char4, char5, char6)

	var num1 uint8 = char6 // why no need type cast bcz byte and uint8 are same

	var num2 int32 = char1

	var num3 int64 = int64(char1)

	println(num1, num2, num3)

	var char7 rune = char1 + char2 + char5 // can perform arth operation on rune or any char bcz they are considered as just numbers

	println(char7)

	char4++ //
	char4++
	char4++
	// char4 = char4 + 1
	// char4 += 1

	// char4-- // decre

	println(char4)

	// to see the number as a char , need to type cast to string

	println(string(char4))

	num1 = 65

	var ch1 rune = 'A'

	println(num1, string(num1))
	println(ch1, string(ch1))

	var num4 int = 19990
	str1 := string(num4) // "19990"

	println(str1)

	str2 := strconv.Itoa(num4) // This would convert the string as digits, can use fmt.Sprint
	//str2 := fmt.Sprint(num4)
	fmt.Println(str2, "-->", reflect.TypeOf(str2))

	// add := calc(10, 20)

	// println(add)

	// add = calc(101, 202)

	// println(add)

	a, b, c, d, e := calc(20, 9)

	println(a, b, c, d, e)

	a1, b1, c1, _, _ := calc(20, 9)

	println(a1, b1, c1)

	a2, _, c2, _, _ := calc(20, 9)

	println(a2, c2)

	str3 := "3123123" // convert to int --> it can be converted or it cannot be converted

	num5, _ := strconv.Atoi(str3) /// use blank identifier for an error is not correct in programs

	println(num5)

	str4 := "312123a" // convert to int --> it can be converted or it cannot be converted

	num6, err := strconv.Atoi(str4) /// use blank identifier for an error is not correct in programs

	if err != nil { // there is an error
		println(">>>>>>", err.Error())
	} else {
		println(num6)
	}

	str5 := "754645.423"

	float1, err := strconv.ParseFloat(str5, 64)

	if err != nil {
		println(">>>", err.Error())
	} else {
		println(float1)
	}

	// any type can be convertd to string

	str6 := fmt.Sprint(3312, true, 'A', 2312.12312)
	//str6 := fmt.Sprintf("%d%v%d%f", 3312, true, 'A', 2312.12312)
	println(str6)
}

// rune -> char --> int32
// byte

// func calc(a int, b int) int {
// 	return a + b
// }

// func calc(a int, b int) (int, int, int, int) {
// 	return a + b, a - b, a * b, a / b
// }

func calc(a int, b int) (add int, sub int, mul int, div int, mod int) {
	add = a + b
	sub = a - b
	mul = a * b
	div = a / b
	mod = a % b
	return add, sub, mul, div, mod
}
