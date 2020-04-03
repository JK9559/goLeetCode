package main

/*
有两种解法：
*/

// 从左侧开始比较，这种方法可能会遇到4, 5, 6, 7, 0, 1, 2的情况
// 在某次划分之后 得到待判断的区间是 0, 1, 2 这时的特殊处理需要考虑左右边界的关系
func findMin1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, h := 0, len(nums)-1
	for l < h {
		if nums[l] < nums[h] {
			return nums[l]
		}
		m := l + (h-l)/2
		// 当m==l的时候
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			h = m
		}
	}
	return nums[l]
}

// 从最右侧开始比较
func findMin(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[m] <= nums[h] {
			h = m
		} else {
			l = m + 1
		}
	}
	return nums[l]
}

func main() {
	findMin([]int{4, 5, 6, 7, 0, 1, 2})
}
