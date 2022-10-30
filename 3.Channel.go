package main

import (
	"fmt"
	"time"
)

func main() {
	// Intro: Channel is mainly used for communicate go routines and function by passing channel as an argument.
	// It's like a pipe through which you can send or receive a message
	// Channels have a type as well so this one will be a string, and we'll only be able to pass messages that are strings any type works though you can even send channels through channels
	ch1 := make(chan string)
	go printValue("dog", ch1)

	msg := <-ch1
	fmt.Println(msg)

	ch2 := make(chan int)
	go countVal(ch2)

	for msg := range ch2 {
		fmt.Println(msg)
	}

}

func printValue(things string, ch chan string) {
	ch <- things
	close(ch)
}

func countVal(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}
