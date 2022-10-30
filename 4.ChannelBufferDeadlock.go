package main

import "fmt"

func main() {
	channelBufferDeadlock(false)
	// We might expect this to work and just output the word "hello" but we're actually going to get deadlock again
	// This is because to send will block until something ready to receive. But the code never progresses to the receive line because we're blocked at sent to make this work we'd need to receive in a separate go routine alternatively we can make buffered channel by giving a capacity when we make the channel.
	// We can fill up a buffered channel without a corresponding receiver, and it won't block until the channel is full.
	channelBufferDeadlock(true)
}

func channelBufferDeadlock(addBuffer bool) {
	ch := make(chan string)
	if addBuffer {
		ch = make(chan string, 2)
	}
	ch <- "hello"

	msg := <-ch
	fmt.Println(msg)
}
