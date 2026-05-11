package main

// channel is used to communicate between goroutines
func main() {

	var ch chan int // this is a a chan which is  declared but not instantiated, so it is nil

	if ch == nil {
		println("nil channel")
		ch = make(chan int) // this is an unbuffered channel, did not mention the size
	}

	go func() {
		// to send a data
		ch <- 100
	}()

	v := <-ch // receive  data
	println(v)

}

// chan is a kind of a queue
// there is unbuffered and buffered channels
// generally two or more goroutines communicate using channels.
// So , there is a sender --> sends the data and there is a receiver that receives the data
// sender is a goroutine and receiver is also a gorotuine

// the sender is blocked , until the receiver receives the data,
// the receiver is blocked, until the sender sends the data
