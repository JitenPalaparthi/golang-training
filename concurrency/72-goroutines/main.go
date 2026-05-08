package main

import (
	"runtime"
	"time"
)

func main() {

	go Even("even-1", 10) // concurrent execution
	go Even("even-2", 20) // concurrent execution
	go Even("even-3", 15) // concurrent execution

	//Even("even-normal", 10) // sequensial execution
	runtime.Goexit() // 1.Exit the goroutine 2. Before going exit, make sure that all other goroutines those are called inside the caller goroutine to complete their execution
	// 3. but while using in main, it would crash the application after completing all other goroutines execution
}

func Even(name string, r int) {
	for i := 1; i <= r; i++ {
		if i%2 == 0 {
			time.Sleep(time.Microsecond * 10)
			println(name, "-->", i)
		}
	}
}
