package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// A WaitGroup waits for a collection of goroutines to finish.
	//A WaitGroup must not be copied after first use.
	var wg sync.WaitGroup

	//Normal function call
	countNum("dog")

	//The main goroutine calls Add to set the number of goroutines to wait for.
	wg.Add(1)
	go func() {
		//Then each of the goroutines runs and calls Done when finished.
		defer wg.Done()
		countNum("cat")
	}()

	//Normal function call
	countNum("rat")

	// Go routine
	wg.Add(1)
	go printVal("cheese", &wg)

	//At the same time, Wait can be used to block until all goroutines have finished.
	wg.Wait()
}

func countNum(things string) {
	for i := 0; i < 5; i++ {
		fmt.Println(things)
		time.Sleep(time.Millisecond * 100)
	}
}

func printVal(things string, wg *sync.WaitGroup) {
	defer wg.Done()
	countNum(things)
}
