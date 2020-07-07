// We can specify if a channel is meant to only send or receive values
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
	// The below won't work. Will throw the following error:
	// invalid operation: <-pings (receive from send-only type chan<- string)
	// // fmt.Println(<-pings)
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	// The below won't work. Will throw the following error:
	// invalid operation: pings <- msg (send to receive-only type <-chan string)
	// // pings <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
