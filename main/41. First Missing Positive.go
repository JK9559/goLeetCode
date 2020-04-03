package main

import "fmt"

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] < len(nums) && nums[nums[i]-1] != nums[i] {
			fmt.Println(i, nums[i])
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
		fmt.Println(nums)
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{-1, 4, 2, 1, 9, 10}))
}
