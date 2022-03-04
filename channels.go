package main

import "fmt"

// Channels
// How concurrent thread connects to goroutines
// Unbuffered by default (only accepts sends)
func main() {
	// Initiate a new channel to store the value
	message := make(chan string)

	go func() {
		// channel <- value will send the value to the channel
		message <- "ping"
	}()

	// <-channel recieves the value from the channel
	msg := <-message
	fmt.Println(msg)

	// Channels can be buffered
	// Thy can accept a limited number of values without a reciever
	bMsg := make(chan string, 2)

	bMsg <- "buffered"
	bMsg <- "channel"

	fmt.Println(<-bMsg)
	fmt.Println(<-bMsg)
	// You need to follow the buffered limit or run into a deadlock issue
	// fmt.Println(<-bMsg)
}
