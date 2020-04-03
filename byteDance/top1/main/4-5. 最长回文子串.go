package main

import (
	"fmt"
	"time"
)

// dp[i][j]表示从字符串s的i到j的字符串是否为回文串,初始化状态为 每个单个字符为回文 连续两个相同字符为回文
// dp[i][j]的状态转移方程为 如果dp[i+1][j-1]为1 那么如果s[i]==s[j] 那么dp[i][j]=1

func longestPalindrome(s string) string {
	lenStr := len(s)
	var dp [][]int
	if 0 == len(s) {
		return ""
	}
	if 1 == len(s) {
		return s
	}
	if 2 == len(s) {
		if s[0] == s[1] {
			return s
		} else {
			return s[0:1]
		}
	}
	keyi, keyj, res := 0, 0, 0
	t1 := time.Now()
	dp = make([][]int, lenStr)
	for i := 0; i < lenStr; i++ {
		dp[i] = make([]int, lenStr)
		dp[i][i] = 1
		if i+1 < lenStr {
			if s[i] == s[i+1] {
				dp[i][i+1] = 1
			}
		}
	}
	for i := 0; i+1 < lenStr; i++ {
		dp[i+1][i] = dp[i][i+1]
	}
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
	t1 = time.Now()
	for i := lenStr - 1; i >= 0; i-- {
		for j := i + 1; j < lenStr; j++ {
			if dp[i+1][j-1] == 1 && s[i] == s[j] {
				dp[i][j] = 1
				if j-i+1 > res {
					keyi = i
					keyj = j
					res = j - i + 1
				}
			}
		}
	}
	t2 = time.Now()
	fmt.Println(t2.Sub(t1))
	return s[keyi : keyj+1]
}

func P05() {
	s := longestPalindrome("abc")
	fmt.Println(s)

}
