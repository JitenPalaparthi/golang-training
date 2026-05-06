package main

import (
	"errors"
	"math/rand/v2"
)

func main() {

	m1 := make(map[string]any)

	m1["add"] = func(i1, i2 int) int {
		return i1 + i2
	}

	m1["sub"] = sub

	mul := func(a, b int) int {
		return a * b
	}

	m1["mul"] = mul

	var fn1 Fn = func(i1, i2 int) int {
		return i1 / i2
	}

	m1["div"] = fn1

	m1["greet"] = func() {
		println("Hello World!")
	}
	m1["done"] = false

	m1["max"] = func(slice []int) (int, error) {
		if slice == nil {
			return 0, errors.New("nil slice")
		}
		if len(slice) > 0 {
			max := slice[0]

			for _, v := range slice {
				if v > max {
					max = v
				}
			}
			return max, nil
		}
		return 0, errors.New("empty slice")
	}

	for k, v := range m1 {

		switch f := v.(type) {
		case bool:
			println(k, "-->", f)
		case func():
			println("executing-->", k)
			f() // just execute

		case func(int, int) int:
			a, b := 20, 10
			println("executing-->", k)
			r := f(a, b)
			println("Result:", r)

		case Fn:
			a, b := 20, 10
			println("executing-->", k)
			r := f(a, b)
			println("Result:", r)

		case func([]int) (int, error):
			var slice []int

			for i := 0; i < 10; i++ {
				slice = append(slice, rand.IntN(100))
			}

			println("executing-->", k)

			max, err := f(slice)

			if err != nil {
				println(err.Error())
			} else {
				println("max:", max)
			}

		default:
			println("key-->", k)
			println("not found the execution type or a value")

		}
		println()
	}

}

// named function
func sub(a, b int) int {
	return a - b
}

type Fn func(int, int) int
