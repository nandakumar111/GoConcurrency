package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	go countIncrement("first", &wg)
	go countIncrement("second", &wg)
	go countIncrement("third", &wg)
	go countIncrement("fourth", &wg)
	go countIncrement("fifth", &wg)

	wg.Wait()
	fmt.Println(counter)

	// This program will run and give the result.
	// But if we check -race in this program, it will show the warning.
	// Because all the go routine function will run simultaneously.
	//counter++ will increment that address of the value.
	// If all the function do parallel that time race condition occur.

	// To check race : go run -race <filename>.go
}

func countIncrement(things string, wg *sync.WaitGroup) {
	fmt.Println(things)
	counter++
	time.Sleep(time.Millisecond * 200)
	wg.Done()
}
