package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it is received from queue
	// Terminates after 2 elements because the channel is closed
	// It is possible to close a non-empty channel
	// It is possible to receive remaining values from closed channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
