package main

// Kadane算法-计算连续子串和时 是dp的一个应用
func maxSubArray(nums []int) int {
	res, curMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMax = max03(nums[i]+curMax, nums[i])
		res = max03(curMax, res)
	}
	return res
}

func max03(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
