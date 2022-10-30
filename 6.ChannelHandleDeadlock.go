package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
		close(ch)
	}()

	for {
		msg, open := <-ch
		if !open {
			break
		}
		fmt.Println(msg)
	}

}
