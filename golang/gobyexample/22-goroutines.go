// A `goroutine` is a lightweight thread of execution
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		// time.Sleep(time.Second)
	}
}

func main() {

	f("direct")

	// Invokes function in a goroutine
	go f("goroutine")

	// You can start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
