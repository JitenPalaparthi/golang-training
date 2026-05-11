package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)

	ch := make(chan int, 5) // buffered channel

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 10)
		v := <-ch // receive  data
		println(v)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println("Sender sending the value", time.Now())
		ch <- 100
		fmt.Println("Sender is not waiting for the receiver to receive", time.Now())
		wg.Done()
	}()

	wg.Wait()

}

// on buffered channels , sender and receiver are not blocked until the buffer is full
