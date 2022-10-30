package main

import (
	"fmt"
	"strconv"
)

func main() {
	oddChannel := make(chan string)
	evenChannel := make(chan string)

	go func(ch chan string, start int, end int) {
		for i := start; i <= end; i++ {
			if i%2 == 1 {
				ch <- strconv.Itoa(i)
			}
		}
	}(oddChannel, 0, 10)

	go func(ch chan string, start int, end int) {
		for i := start; i <= end; i++ {
			if i%2 == 0 {
				ch <- strconv.Itoa(i)
			}
		}
	}(evenChannel, 0, 10)

	for {
		select {
		case odd := <-oddChannel:
			fmt.Println(odd)
		case even := <-evenChannel:
			fmt.Println(even)
		}
	}
	// Note: We will get deadlock here. because we didn't close this channel. if we close this channel, it will go to infinite
}
