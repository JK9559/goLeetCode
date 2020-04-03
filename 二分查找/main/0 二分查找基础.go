package main

import "fmt"

func binarySearch(ary []int, target int) int {
	l, h := 0, len(ary)-1

	for l <= h {
		m := l + (h-l)/2
		if ary[m] > target {
			h = m - 1
		} else if ary[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3))
}
