package main

import "time"

func main() {

	go func() {
		println("Hello World! Goroutine")
	}() // this func run concurrently

	go Even("even-1", 100)
	go Even("even-2", 100)

	println("Hello World by Main")

	time.Sleep(time.Millisecond * 1000) // wait for 100 milliseconds, block here, this is not a solution

} // main exits

func Even(name string, r int) {
	for i := 1; i <= r; i++ {
		if i%2 == 0 {
			time.Sleep(time.Microsecond * 10)
			println(name, "-->", i)
		}
	}
}

// main is also a goroutine, though we dont create, but the goruntime executes it as a goroutine
// to create a goroutine , use a keyword called go
// no goroutine would wait for other goroutine to complete its execution
// the order of execution is not guaranteed

// concurrency is not parallelism

// executing multiple things
// parallel means , executing multiple things at once

// some hardware cores --> 8 cores
// Multiplexing --> Run many number of systems threads on a few number of hardware cores
// os has a scheduler -> it schedules to run each os thread with a stipulated time -> ex 10 ms

// instead of depending on os scheduler, it depends on its own scheduler
// Go runtime -> its own scheduler
// We dont create threads in go, rather we create go routines

// Go routines are very small, agile, can create 1000s of goroutines seamlessly, go runtime handles them
// generaly each thread has to be created by OS but in Go , goroutines are created by goruntime
// goroutine is a green thread, a simple , thread managed by Goruntime
