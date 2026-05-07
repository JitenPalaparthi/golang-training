package main

func main() {
	defer println("End of main-1")

	defer println("End of main-2")

	defer func() { // main.func1
		println("End of main-3")
	}()

	func() { // main.func1
		defer println("end of main.func1")
		func() { // main.func1.func2
			defer println("End of main-4")
		}()
		println("start of main.func1")
	}()

	println("Start of Main")

}

// defer -> defer defers the execution to the end of caller
// for every caller, defer maintains the stack,  first in last out or last in first out
