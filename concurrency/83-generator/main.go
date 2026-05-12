package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := Generator("gen-1", 10)
	ch2 := Generator("gen-2", 10)
	ch3 := Generator("gen-3", 10)

	<-Receive(ch1)
	<-Receive(ch2)
	<-Receive(ch3)

	// <-sig
	// <-sig
	// <-sig // This blocks the main until the signal is received

}

func Generator(name string, r int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= r; i++ {
			if i%2 == 0 {
				time.Sleep(time.Millisecond * 100)
				ch <- fmt.Sprint(name, "--->", i)
			}
		}
		close(ch)
	}()
	return ch
}

func Receive(ch <-chan string) <-chan struct{} { // receiver
	sig := make(chan struct{})
	go func() {
		for v := range ch { // The range loop iterates until the channel is not closed
			println(v)
		}
		sig <- struct{}{}
	}()
	return sig
}

// func Receive(ch <-chan string, sig chan<- struct{}) { // receiver
// 	for v := range ch { // The range loop iterates until the channel is not closed
// 		println(v)
// 	}
// 	sig <- struct{}{}
// }

// range loop on channels givves only data received unlike arrays/slices/maps
