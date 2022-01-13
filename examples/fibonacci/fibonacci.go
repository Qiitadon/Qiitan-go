package main

import "fmt"

func fib(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}

	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println(fib(35))
}

// Output: 9227465
