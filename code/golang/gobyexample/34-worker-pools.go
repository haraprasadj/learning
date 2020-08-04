package main

import (
	"fmt"
	"time"
)

// The worker, of which we'll run concurrent instances.
// The workers receive jobs on `jobs` channel and send results on `results` channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- 2 * j
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start up 3 workers, initially blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs and close the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// Finally collect results
		<-results
	}
}
