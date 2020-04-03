package main

import "fmt"

/*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
给出两个整数 x 和 y，计算它们之间的汉明距离。
*/

func hammingDistance(x int, y int) int {
	z := x ^ y
	cnt := 0
	for z != 0 {
		// 去除z中二进制位中最低的那一位
		z &= z - 1
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
