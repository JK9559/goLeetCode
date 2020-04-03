package main

func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0

	for left < right {
		res = max02(min02(height[left], height[right])*(right-left), res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min02(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max02(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
