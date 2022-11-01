package main

import "fmt"

var ch = make(chan string)

func main() {
	go ping("jj:")
	ch <- "Bye!!"
	fmt.Println(<-ch)
}

func ping(things string) {
	fmt.Println(things, <-ch)
	ch <- "Have a nice day!!"
}
