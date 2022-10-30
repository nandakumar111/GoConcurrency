package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	incrVal int32
	wg      sync.WaitGroup
)

func main() {
	wg.Add(3)

	go countInc("one")
	go countInc("three")
	go countInc("four")

	wg.Wait()
	fmt.Println(incrVal)
}

func countInc(things string) {
	defer wg.Done()
	fmt.Println(things)
	//counter++ // Normal value increment
	// Using atomic function
	atomic.AddInt32(&incrVal, 1)
	time.Sleep(time.Millisecond * 200)
}
