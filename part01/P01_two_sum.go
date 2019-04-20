package part01

import "fmt"

func twoSum(nums []int, target int) []int {
	rstMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rstMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		tar := target - nums[i]
		tarJ := rstMap[tar]
		if tarJ > i {
			return []int{i,tarJ}
		}
	}
	return []int{}
}

func P01() {
	req := []int{-1, -2, -3, -4, -5}
	target := -8
	rst := twoSum(req, target)
	fmt.Println(rst)
}
