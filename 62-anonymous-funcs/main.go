package main

import (
	"errors"
	"strings"
)

func main() {

	r1 := calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	println("sum:", r1)

	r2 := calc(20, 10, sub) // passing a named function as an argument

	println("sub:", r2)

	mul := func(a, b int) int {
		return a * b
	}

	r3 := calc(10, 5, mul)

	println("mul:", r3)

	f1, err := calcr("addition")

	if err == nil {
		r := f1(10, 20)
		println("result:", r)
	} else {
		println(err.Error())
	}

	f2, err := calcr("mod")

	if err != nil {
		println(err.Error())
	} else {
		r := f2(10, 2)
		println("result:", r)
	}
}

// functions can also be used like a variable

func calc(a int, b int, fn func(int, int) int) int {
	return fn(a, b)
}

// named function
func sub(a, b int) int {
	return a - b
}

func calcr(ops string) (func(int, int) int, error) {
	ops = strings.ToLower(ops)
	switch ops {
	case "add", "addition", "sum":
		return func(i1, i2 int) int {
			return i1 + i2
		}, nil
	case "sub", "substract":
		return func(i1, i2 int) int {
			return i1 - i2
		}, nil
	case "mul", "multiplication", "multiply":
		return func(i1, i2 int) int {
			return i1 * i2
		}, nil
	case "div", "divide", "division":
		return func(i1, i2 int) int {
			return i1 / i2
		}, nil
	default:
		return nil, errors.New("no ops")
	}
}
