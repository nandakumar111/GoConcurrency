package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	incr      int32
	waitGroup sync.WaitGroup
	mutex     sync.Mutex
)

func main() {
	waitGroup.Add(3)

	go countIncr("one")
	go countIncr("three")
	go countIncr("four")

	waitGroup.Wait()
	fmt.Println(incr)
}

func countIncr(things string) {
	defer waitGroup.Done()
	fmt.Println(things)
	mutex.Lock()
	{
		incr++
	}
	mutex.Unlock()
	time.Sleep(time.Millisecond * 200)
}
