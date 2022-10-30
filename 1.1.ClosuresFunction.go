package main

import "fmt"

func main() {
	// Closures are a special case of anonymous functions.
	// Closures are anonymous functions which access the variables defined outside the body of the function.

	a, b := 2, 3
	func() {
		fmt.Println(a + b)
	}()

	r := 4.2
	area := func(radius float64) float64 {
		return 3.14 * r * r
	}(r)
	fmt.Println(area)

}
