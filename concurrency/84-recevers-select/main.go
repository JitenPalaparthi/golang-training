package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := Generator("gen-1", 10)
	ch2 := Generator("gen-2", 10)
	ch3 := Generator("gen-3", 10)

	<-ReceiveSelect(ch1, ch2, ch3)

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

func ReceiveSelect(ch1 <-chan string, ch2 <-chan string, ch3 <-chan string) <-chan struct{} {
	sig := make(chan struct{})
	done := 0
	go func() {
		for {
			if done == 3 {
				sig <- struct{}{}
				break
			}
			select {
			case v, ok := <-ch1:
				if ok {
					println(v)
				} else {
					done++
				}
			case v, ok := <-ch2:
				if ok {
					println(v)
				} else {
					done++
				}

			case v, ok := <-ch3:
				if ok {
					println(v)
				} else {
					done++
				}
			}
		}
	}()
	return sig
}
