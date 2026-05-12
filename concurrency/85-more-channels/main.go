package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := Generator("gen-1", 10)
	ch2 := Generator("gen-2", 10)
	ch3 := Generator("gen-3", 10)
	ch4 := Generator("gen-4", 10)
	ch5 := Generator("gen-5", 10)
	ch6 := Generator("gen-6", 10)

	<-Receive(ch1, ch2, ch3, ch4, ch5, ch6)

}

func Generator(name string, r int) chan string {
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

func Receive(chs ...chan string) chan struct{} {
	wg := new(sync.WaitGroup)
	sig := make(chan struct{})
	wg.Add(len(chs))
	wg.Add(1)
	wg.Add(1)
	go func() {
		wg.Wait()
		sig <- struct{}{}
	}()

	go func() {
		go func() {
			for _, ch := range chs {
				go func() {
					for v := range ch {
						println(v)
					}
					wg.Done()
				}()
			}
			wg.Done()
		}()
		wg.Done()
	}()
	return sig
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
