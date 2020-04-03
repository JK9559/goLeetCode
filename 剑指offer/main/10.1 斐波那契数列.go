package main

import "fmt"

var fib []int

func Fib40() {
	fib = make([]int, 40)
	fib[1] = 1
	for i := 2; i < 40; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
}

func Fibonacci(n int) int {
	return fib[n]
}

func main() {
	Fib40()
	for i := 0; i < 40; i++ {
		fmt.Printf("%d ", Fibonacci(i))
	}
}
