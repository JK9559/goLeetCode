package main

import "fmt"

/*
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。
*/

// 首先遍历数组 异或 得到最终结果两个数的异或结果
// 得到两个数不同的比特位 然后用异或结果 区分整个数组 得到两个数
func singleNumberIII(nums []int) []int {
	dif := 0
	for i := 0; i < len(nums); i++ {
		dif ^= nums[i]
	}
	// 取异或结果的最右位
	dif &= -dif
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		if dif&nums[i] == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

func main() {
	fmt.Println(singleNumberIII([]int{1, 1, 2, 2, 3, 4}))
}
