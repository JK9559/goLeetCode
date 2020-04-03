package main

import "fmt"

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
*/

/*
使用二分法 可以区分旋转数组和递增数组
此时问题的关键在于确定对半分得到的两个数组哪一个是旋转数组，哪一个是非递减数组。我们很容易知道非递减数组的第一个元素一定小于等于最后一个元素。

通过修改二分查找算法进行求解（l 代表 low，m 代表 mid，h 代表 high）：

当 nums[m] <= nums[h] 时，表示 [m, h] 区间内的数组是非递减数组，[l, m] 区间内的数组是旋转数组，此时令 h = m；
否则 [m + 1, h] 区间内的数组是旋转数组，令 l = m + 1。
*/

func minNumberInRotateArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2
		if nums[l] < nums[m] {
			l = m + 1
		} else if nums[l] == nums[m] && nums[m] == nums[h] {
			for i := l; i < h; i++ {
				if nums[i] > nums[i+1] {
					return nums[i+1]
				}
			}
			return nums[l]
		} else {
			h = m
		}
	}
	return nums[l]
}

func main() {
	nums := []int{1, 1, 1, 1}

	fmt.Println(minNumberInRotateArray(nums))
}
