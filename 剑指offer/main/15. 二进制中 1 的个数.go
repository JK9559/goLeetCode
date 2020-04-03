package main

import "fmt"

func NumberOf1(n int) int {
	cnt := 0
	for n != 0 {
		cnt++
		n &= n - 1
	}
	return cnt
}

func main() {
	fmt.Println(NumberOf1(5))
	fmt.Printf("%b\n", 5)
}
