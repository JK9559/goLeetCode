package main

import (
	"fmt"
	"math"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res <<= 1
		res |= head.Val
		head = head.Next
	}
	return res
}

/*
我们定义「顺次数」为：每一位上的数字都比前一位上的数字大 1 的整数。

请你返回由 [low, high] 范围内所有顺次数组成的 有序 列表（从小到大排序）
*/
func sequentialDigits(low int, high int) []int {
	lowS := strconv.Itoa(low)
	highS := strconv.Itoa(high)
	ans := low
	res := make([]int, 0)

	for key := low / int(math.Pow10(len(lowS)-1)); ; {
		fmt.Println(key)
		j := key
		ans1 := 0
		for ans1 = key; len(strconv.Itoa(ans1)) < len(lowS); {
			ans1 = ans1 * 10
			j++
			ans1 += j
		}
		fmt.Println(ans1)
		fmt.Println(int(math.Pow10(len(lowS))))
		if ans1 < int(math.Pow10(len(lowS))) && ans1 >= low {
			res = append(res, ans1)
		} else {
			break
		}
	}

	for i := len(lowS) + 1; ans <= high && i <= len(highS); i++ {
		for key := 1; key <= 10-i; key++ {
			j := key
			for ans = key; len(strconv.Itoa(ans)) < i; {
				ans = ans * 10
				j++
				ans += j
			}
			if ans <= high {
				res = append(res, ans)
			}
		}
	}
	return res
}

func pow(x, y int) {

}

/*
5285. 元素和小于等于阈值的正方形的最大边长 显示英文描述
用户通过次数274
用户尝试次数430
通过次数276
提交次数760
题目难度Medium
给你一个大小为 m x n 的矩阵 mat 和一个整数阈值 threshold。

请你返回元素总和小于或等于阈值的正方形区域的最大边长；如果没有这样的正方形区域，则返回 0 。
*/
func maxSideLength(mat [][]int, threshold int) int {

}

/*
给你一个 m * n 的网格，其中每个单元格不是 0（空）就是 1（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。

如果您 最多 可以消除 k 个障碍物，请找出从左上角 (0, 0) 到右下角 (m-1, n-1) 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 -1。
*/
func shortestPath(grid [][]int, k int) int {

}

func main() {
	fmt.Println(sequentialDigits(58, 155))
}
