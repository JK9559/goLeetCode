package main

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candies[i] = 1
	}

	// 从左往右遍历 只有当 右侧大于左侧时 更新右侧得分为左侧得分+1
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右侧往左侧遍历 只有当 左侧大于右侧 并且 左侧得分小于或者等于右侧得分时 才将左侧得分更新为右侧+1
	sum := candies[len(ratings)-1]
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] && candies[i-1] <= candies[i] {
			candies[i-1] = candies[i] + 1
		}
		sum += candies[i-1]
	}
	return sum
}

func main() {

}
