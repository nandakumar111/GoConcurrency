package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			val := strconv.Itoa(i)
			if i%2 == 0 {
				ch <- val + ": Even"
			} else {
				ch <- val + ": Odd"
			}
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
