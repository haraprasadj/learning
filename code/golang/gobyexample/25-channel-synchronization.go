// We can use channels to synchronize execution across goroutines
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we are done
	done <- true
}

func main() {
	done := make(chan bool, 1)

	// Start `worker` goroutine, giving it the channel to notify on
	go worker(done)

	// Block until we receive a notification from the worker on the channel
	<-done
}
