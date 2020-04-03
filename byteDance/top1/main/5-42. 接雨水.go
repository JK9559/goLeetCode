package main

// https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/

// 叠加法
// 理解这个方法 需要看图
func trapByDP(height []int) int {
	if len(height) == 0 {
		return 0
	}

	ans := 0
	max_left := make([]int, len(height))
	max_right := make([]int, len(height))

	max_left[0] = height[0]
	max_right[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		max_left[i] = max01(height[i], max_left[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		max_right[i] = max01(height[i], max_right[i+1])
	}
	for i := 0; i < len(height); i++ {
		ans += min01(max_right[i], max_left[i]) - height[i]
	}
	return ans
}

func max01(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min01(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 双指针法
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	ans, max_left, max_right := 0, height[0], height[len(height)-1]
	left, right := 0, len(height)-1

	for left < right {
		if height[left] < height[right] {
			if max_left > height[left] {
				ans += max_left - height[left]
			} else {
				max_left = height[left]
			}
			left += 1
		} else {
			if max_right > height[right] {
				ans += max_right - height[right]
			} else {
				max_right = height[right]
			}
			right -= 1
		}
	}
	return ans
}

func main() {

}
