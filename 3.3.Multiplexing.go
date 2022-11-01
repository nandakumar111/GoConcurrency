package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring1("ane"), boring1("joe"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both are boring. I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring1(msg string) <-chan string { // Returns receive-only channel of string
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
