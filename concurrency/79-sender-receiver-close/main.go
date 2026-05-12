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
	for {
		v, ok := <-ch // if ok is false that means the channel is cloed
		if ok {
			println(v)
		} else {
			break
		}
	}
	wg.Done()
}
