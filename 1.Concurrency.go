package main

import (
	"fmt"
	"time"
)

func main() {
	//main function won't wait for go routine
	startTime := time.Now()

	//Normal function call
	count("dog")

	// Go routine
	go count("cat")

	//Normal function call
	count("rat")

	// Go routine
	go count("cheese")

	fmt.Println("Time Dif : ", time.Since(startTime))
}

func count(things string) {
	for i := 0; i < 5; i++ {
		fmt.Println(things)
		time.Sleep(time.Millisecond * 100)
	}
}
