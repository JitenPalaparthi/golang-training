package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		Publisher(ch, "pub-1", 2)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Publisher(ch, "pub-2", 2)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		Publisher(ch, "pub-3", 3)
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		println(v)
	}

}

func Publisher(ch chan string, name string, r int) {
	for i := 1; i <= r; i++ {
		ch <- fmt.Sprint("Publisher->", name, "-->", i)
		time.Sleep(time.Millisecond * 100)
	}
	//close(ch) // should never close the channel inside a func/goroutine if the same chan is used by multiple publishers
}
