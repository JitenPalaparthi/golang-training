package main

import (
	"time"
)

func main() {

	ch := make(chan int)

	sig := make(chan struct{})

	go Sender(ch, 10)

	go Receive(ch, sig)

	<-sig // This blocks the main until the signal is received

}

func Sender(ch chan<- int, r int) { // sender
	for i := 1; i <= r; i++ {
		if i%2 == 0 {
			time.Sleep(time.Millisecond * 100)
			ch <- i
		}
	}
	close(ch)
}

func Receive(ch <-chan int, sig chan<- struct{}) { // receiver
	for v := range ch { // The range loop iterates until the channel is not closed
		println(v)
	}
	sig <- struct{}{}
}

// range loop on channels givves only data received unlike arrays/slices/maps
