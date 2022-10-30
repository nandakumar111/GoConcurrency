package main

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	n := 10
	go func(n int) {
		for i := 0; i < n; i++ {
			println(<-ch)
		}
		quit <- true
	}(n)
	fib(ch, quit)
}

func fib(num chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case num <- x:
			x, y = y, x+y
		case <-quit:
			return
		}
	}
}
