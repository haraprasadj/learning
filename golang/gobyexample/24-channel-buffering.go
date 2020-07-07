// By default channels are unbuffered, meaning they will accept sends if there is
// a corresponding receive ready to receive the sent value
package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	messages <- "buffered"
	fmt.Println(<-messages)
	messages <- "channel"
	messages <- "again"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
