package main

import (
	"fmt"
	"time"
)

func main() {

	// Timers represent a single event in the future
	// You can tell a timer how long to wait, and it provides a channel that will be notified at that time
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks the timer's channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Timers can be cancelled before they fire
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Giving the timer2 enough time to fire, to ensure it was in fact stopped
	time.Sleep(2 * time.Second)
}
