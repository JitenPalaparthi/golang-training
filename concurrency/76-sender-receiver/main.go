package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup)

	ch := make(chan int)

	// wg.Add(1)
	// go func() {
	// 	fmt.Println("Receiver is hit", time.Now())
	// 	v := <-ch // receive  data
	// 	println(v)
	// 	fmt.Println("Received a value", time.Now())
	// 	wg.Done()
	// }()
	// wg.Add(1)
	// go func() {
	// 	time.Sleep(time.Second * 10)
	// 	ch <- 100
	// 	wg.Done()
	// }()

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
		fmt.Println("Sender is waiting for the receiver to receive", time.Now())
		wg.Done()
	}()

	wg.Wait()

}
