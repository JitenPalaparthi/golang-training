package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	//wg.Add(6)

	wg.Add(1)
	go Even(wg, "even-1", 1, 10, time.Microsecond*100)

	wg.Add(1)
	go Odd(wg, "odd-1", 1, 10, time.Microsecond*111)

	wg.Add(1)
	go Even(wg, "even-2", 20, 30, time.Microsecond*100)

	wg.Add(1)
	go Odd(wg, "odd-2", 20, 30, time.Microsecond*111)

	wg.Add(1)
	go Odd(wg, "odd-3", 200, 300, time.Microsecond*1)

	// runtime.Goexit() not a solution
	// time.Slee(time.Second*2) not a solution

	wg.Add(1)

	go func() {
		PrimeNumbers("Prime-1", 10, 200, time.Microsecond*100)
		wg.Done()
	}()
	wg.Wait() // Wait until the state become 0

	println("The wait is over, so exisitng main gracefully")

}

func Even(wg *sync.WaitGroup, name string, s, e int, stime time.Duration) {
	for i := s; i <= e; i++ {
		if i%2 == 0 {
			time.Sleep(stime)
			println(name, "-->", i)
		}
	}
	wg.Done() // state=state-1
}

func Odd(wg *sync.WaitGroup, name string, s, e int, stime time.Duration) {
	for i := s; i <= e; i++ {
		if i%2 != 0 {
			time.Sleep(stime)
			println(name, "-->", i)
		}
	}
	wg.Done() // state=state-1
}

func PrimeNumbers(name string, s, r int, stime time.Duration) {
	for i := s; i < r; i++ {
		time.Sleep(stime)
		isPrint := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrint = false
				break
			}
		}
		if isPrint {
			println(name, "-->", i)
		}
	}
}

// WaitGroup --> State/Counter --> 4
