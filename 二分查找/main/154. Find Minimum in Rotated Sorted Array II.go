package main

import "fmt"

func findMinII(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[m] < nums[h] {
			h = m
		} else if nums[m] == nums[h] {
			if nums[l] > nums[m] {
				h = m
			} else {
				//1.
				//for i := l; i < h; i++ {
				//	if nums[i] > nums[i+1] {
				//		return nums[i+1]
				//	}
				//}
				//return nums[l]
				//2.
				h--
			}

		} else {
			l = m + 1
		}
	}
	return nums[l]
}

func main() {
	fmt.Println(findMinII([]int{1, 1, 1, 1, 3, 3, 1, 1, 1, 1}))
}
