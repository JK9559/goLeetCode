package main

import "fmt"

// 用一个map保存值和对应的下标，可以在知道目标之后 很快的找到相应的下标

func twoSum(nums []int, target int) []int {
	rstMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rstMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		tar := target - nums[i]
		tarJ := rstMap[tar]
		if tarJ > i {
			return []int{i, tarJ}
		}
	}
	return []int{}
}

func main() {
	req := []int{-1, -2, -3, -4, -5}
	target := -8
	rst := twoSum(req, target)
	fmt.Println(rst)
}
