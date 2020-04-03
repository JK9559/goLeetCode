package main

import (
	"fmt"
)

// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	minIndex := findSmallest(nums)
	if target == nums[minIndex] {
		return minIndex
	}
	if target < nums[minIndex] || (minIndex > 0 && nums[minIndex-1] < target) || (minIndex == 0 && nums[len(nums)-1] < target) {
		return -1
	}
	if nums[0] < nums[len(nums)-1] {
		return binFind(nums, 0, len(nums)-1, target)
	} else if nums[minIndex] <= target && nums[len(nums)-1] >= target {
		return binFind(nums, minIndex, len(nums)-1, target)
	} else if nums[0] <= target && nums[minIndex-1] >= target {
		return binFind(nums, 0, minIndex-1, target)
	} else {
		return -1
	}
}

// 二分查找
func binFind(nums []int, l, h, target int) int {
	for l <= h {
		m := l + (h-l)/2
		if nums[m] > target {
			h = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

// 旋转数组 寻找最小值(旋转点)
func findSmallest(nums []int) int {
	l, h := 0, len(nums)-1
	if nums[l] < nums[h] {
		return 0
	}
	for l <= h {
		m := l + (h-l)/2
		if m+1 >= len(nums) {
			return 0
		}
		if nums[m] > nums[m+1] {
			return m + 1
		} else {
			if nums[m] < nums[l] {
				h = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(findSmallest([]int{1, 1, 1, 1, 0}))
}
