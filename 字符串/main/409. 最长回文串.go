package main

import "fmt"

func longestPalindrome(s string) int {
	cnts := make([]int, 256)
	for _, v := range s {
		cnts[v]++
	}
	res := 0
	for _, cnt := range cnts {
		res += cnt / 2 * 2
	}
	if res < len(s) {
		res++
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
