package main

func main() {

	var fn1 Fn1 = func(i1, i2 int) int {
		return i1 + i2
	}

	r := fn1(123, 123)
	println("result:", r)

	var sq Fn2 = func(i int) int {
		return i * i
	}

	r = sq(10)
	println("result:", r)

	sq.Greet()

	var fn3 Fn3 = func(i *int) Fn4 {
		*i = *i * *i
		var fn Fn4 = func() int {
			return *i * *i
		}
		return fn
	}

	// if fn3 == nil {
	// 	println("nil function")
	// }

	num := new(10)

	fn4 := fn3(num)

	println("num:", *num)

	r = fn4()
	println("result:", r)

}

type Fn1 func(int, int) int

type Fn2 func(int) int

func (f Fn2) Greet() {
	println("Hello,Welcome to Fn2 type calling Greet method")
}

type Fn3 func(*int) Fn4

type Fn4 func() int
