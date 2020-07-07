// Go's `select` lets you wait on multiple channel operations
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	fmt.Println("started at", time.Now().UTC())
	for i := 0; i < 2; i++ {
		// Use `select` to await both values and print each one as it arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1, "at", time.Now().UTC())
		case msg2 := <-c2:
			fmt.Println("received", msg2, "at", time.Now().UTC())
		}
		// Total time taken is ~2 seconds only since both goroutines execute concurrently
	}
}
