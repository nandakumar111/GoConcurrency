package main

import "fmt"

func main() {
	//A defer statement defers the execution of a function until the surrounding function returns.
	//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println("Main Function End..!!")

	n := 1
	for {
		fmt.Println(n)
		n++
		if n > 10 {
			break
		}
	}
	// Deferred functions are executed in LIFO order, so the above code prints: 1.numbers 2.Loop End..!! 3.Main Function End..!!
	defer fmt.Println("Loop End..!!")
}
