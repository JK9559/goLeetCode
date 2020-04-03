package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。
*/

func searchRange(nums []int, target int) []int {
	first := findFirstEle(nums, target)
	last := findFirstEle(nums, target+1) - 1
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	} else {
		return []int{first, last}
	}
}

// 二分查找 找到最先出现的元素
func findFirstEle(nums []int, target int) int {
	l, h := 0, len(nums)
	for l < h {
		m := l + (h-l)/2
		if target <= nums[m] {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(findFirstEle([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(findFirstEle([]int{5, 7, 7, 8, 8, 10}, 9))
}
