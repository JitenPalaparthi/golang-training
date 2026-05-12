package main

import (
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(1)
	go Even(wg, ch, 100)

	wg.Add(1)
	go Receive(wg, ch)

	wg.Wait()
}

func Even(wg *sync.WaitGroup, ch chan int, r int) {
	for i := 1; i <= r; i++ {
		if i%2 == 0 {
			time.Sleep(time.Millisecond * 100)
			ch <- i
		}
	}
	close(ch)
	wg.Done()
}

func Receive(wg *sync.WaitGroup, ch chan int) {
	for v := range ch { // The range loop iterates until the channel is not closed
		println(v)
	}
	wg.Done()
}

// range loop on channels givves only data received unlike arrays/slices/maps
