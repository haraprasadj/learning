// Closing a channel indicates that no more values will be sent on it
package main

import (
	"fmt"
)

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// In this special 2-value form of receive, the `more` value will be false
			// if `jobs` has been closed and all values on the channel have been received
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
