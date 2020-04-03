package main

import "fmt"

/*
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
*/

// 0 ^ num[0] ^ 1 ^ nums[1] ...
// 两个相同的数异或的结果为 0，对所有数进行异或操作，最后的结果就是单独出现的那个数
func missingNumber(nums []int) int {
	ret := 0
	for i := 0; i < len(nums); i++ {
		ret = ret ^ i ^ nums[i]
	}
	return ret ^ len(nums)
}

func main() {
	fmt.Println(missingNumber([]int{1, 2}))
}
