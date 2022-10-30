package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			ch1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			ch2 <- "Every 2s"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

	// Note: We will get deadlock here. because we didn't close this channel. if we close this channel, it will go to infinite

}
