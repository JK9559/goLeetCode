package main

/*
给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。

注意：1 ≤ k ≤ n ≤ 10^9。
*/

func findKthNumber(n int, k int) int {
	cur := 1
	// 因为cur从1 开始 已经占据了1个位置 所以 k要减1
	k = k - 1
	for k > 0 {
		// 获取两个相邻节点之间 要经过多少步 才能到
		step := getStep(n, cur, cur+1)
		// 如果需求的值 大于 step 说明 已经跨过了这个点 所以 cur向相邻的节点移动 cur++
		// 反之 如果 step 大于 k 说明并未跨过cur 所以将cur向下移动 cur*=10
		if k >= step {
			k -= step
			cur += 1
		} else {
			cur *= 10
			k -= 1
		}
	}
	return cur
}

// 计算两节点之间 经过多少步 才能到
func getStep(n int, from int, to int) int {
	step := 0
	// 只有当from小于或者等于最大数字n 时
	for from <= n {
		// 为了计算当n在 from 和 to之间时的步数
		step += min03(n+1, to) - from
		// 乘以10 是为了移动到下一层
		from *= 10
		to *= 10
	}
	return step
}

func min03(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {

}
