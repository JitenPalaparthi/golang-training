package main

import "sync"

func main() {

	ch := Generator(100)
	<-Workers(10, ch)

}

func Generator(r int) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= r; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Workers(workers int, ch chan int) chan struct{} {
	sig := make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(workers)
	wg.Add(1)

	go func() {
		wg.Wait()
		sig <- struct{}{}
	}()
	go func() {
		for i := range workers {
			go func(work int) {
				for v := range ch {
					if v%2 == 0 {
						println("Even", "Wroker->", work+1, " Data:", v)
					} else {
						println("Odd", "Wroker->", work+1, " Data:", v)
					}
				}
				wg.Done()
			}(i)
		}
		wg.Done()
	}()
	return sig
}
