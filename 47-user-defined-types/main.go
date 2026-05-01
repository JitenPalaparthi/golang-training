package main

func main() {

	var a *A
	var b B

	Greet()

	a.Greet()
	b.Greet()

	(&A{}).Greet()

	(&B{}).Greet()

}

type A struct{}

type B struct{}

func Greet() {
	println("Hello World, I am from a func Greet")
}

func (*A) Greet() {
	println("Hello World, I am from type A")
}

func (*B) Greet() {
	println("Hello World, I am from type B")
}

// can think it like polymorphisam, generally go does not implement poly by design like OOPS programming languages
