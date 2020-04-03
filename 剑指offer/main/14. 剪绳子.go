package main

import "fmt"

/*
把一根绳子剪成多段，并且使得每段的长度乘积最大。
*/

// dp
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			// j * (i-j) 是计算当dp[i] < i 时的情况 如dp[2]=1
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(integerBreak(2))
}
