package main

import "fmt"

func maxProduct(words []string) int {
	n := len(words)
	vals := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(words[i]); j++ {
			vals[i] |= 1 << (words[i][j] - 'a')
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if vals[i]&vals[j] == 0 {
				res = maxInt(res, len(words[i])*len(words[j]))
			}
		}
	}
	return res
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))

}
