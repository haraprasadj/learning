// Use `select` with a `default` clause to implement non-blocking sends, receives
// and even non-blocking multi-way `selects`
package main

import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	// Non-blocking receive
	// If a value is not available on the channel `select` takes the default case
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking send
	// Here, `msg` cannot be sent because the channel has no buffer and no receiver
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// Use multiple cases above `default` to create a non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
