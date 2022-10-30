package main

import "fmt"

func main() {
	//A Higher-Order function is a function that receives a function as an argument or returns the function as output.
	//Higher order functions are functions that operate on other functions, either by taking them as arguments or by returning them.
	fmt.Println(sum(3)(4))
	fmt.Println(sumFive(1)(2)(3)(4)(5))
}

func sum(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// Returning Functions from other Functions

type First func(int) int
type Second func(int) First
type Third func(int) Second
type Fourth func(int) Third

func sumFive(a int) Fourth {
	return func(b int) Third {
		return func(c int) Second {
			return func(d int) First {
				return func(e int) int {
					return a + b + c + d + e
				}
			}
		}
	}
}
