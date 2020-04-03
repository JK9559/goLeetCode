package main

import "fmt"

/*
给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
*/

// n&n-1 去除了最低位的1 dp[n] 表示n的二进制表示有多少1
// dp[n] = dp[n&(n-1)]+1 表示n的二进制1数目 等于 去掉最低位1的数字的二进制1数目 加上 1
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func main() {
	fmt.Println()
}
