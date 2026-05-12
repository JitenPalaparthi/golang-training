package main

import (
	"time"
)

func main() {

	sig := make(chan struct{})

	ch := Generator(10)

	go Receive(ch, sig)

	<-sig // This blocks the main until the signal is received

}

func Generator(r int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= r; i++ {
			if i%2 == 0 {
				time.Sleep(time.Millisecond * 100)
				ch <- i
			}
		}
		close(ch)
	}()
	return ch
}

func Receive(ch <-chan int, sig chan<- struct{}) { // receiver
	for v := range ch { // The range loop iterates until the channel is not closed
		println(v)
	}
	sig <- struct{}{}
}

// range loop on channels givves only data received unlike arrays/slices/maps
