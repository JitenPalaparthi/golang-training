package main

import (
	"sync"
	"time"
)

func main() {

	ch := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				time.Sleep(time.Millisecond * 100)
				ch <- i
			}
		}
		wg.Done()

	}()

	go func() {
		for i := 1; i <= 5; i++ {
			println(<-ch)
		}
		wg.Done()
	}()
	wg.Wait()

}
