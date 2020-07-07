// Channels are the pipes that connect the concurrent goroutines
package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a new channel. Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the <- syntax
	// Here we send "ping" to the `messages` channel, from a goroutine
	go func() {
		messages <- "ping"
		time.Sleep(time.Second)
		messages <- "ping again"
	}()

	// Receive a value from the channel
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
