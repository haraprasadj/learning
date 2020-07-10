package main

import (
	"fmt"
	"time"
)

func main() {

	// Ticker provides a channel that is sent values at the defined time interval
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Tickers can be stopped like timers
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
