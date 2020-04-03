package main

import "fmt"

/*
判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一个格子开始，
每一步可以在矩阵中向上下左右移动一个格子。如果一条路径经过了矩阵中的某一个格子，则该路径不能再进入该格子。
*/

/*
回溯法
使用回溯法（backtracking）进行求解，它是一种暴力搜索方法，通过搜索所有可能的结果来求解问题。
回溯法在一次搜索结束时需要进行回溯（回退），将这一次搜索过程中设置的状态进行清除，从而开始一次新的搜索过程。
*/

func hasPath(array string, rows, cols int, str string) bool {
	// 方向数组 上开始 顺时针
	direct := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	strMatrix := buildMatrix(array, rows, cols)
	strByte := []byte(str)
	visit := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visit[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if isTharOk(strMatrix, i, j, visit, direct, strByte, 0) {
				return true
			}
		}
	}
	return false
}

func isTharOk(matrix [][]byte, r, c int, visit [][]bool, direct [][]int, str []byte, nowPos int) bool {
	if nowPos == len(str)-1 {
		return true
	}
	if r < 0 || r >= len(matrix) || c < 0 || c >= len(matrix[0]) || matrix[r][c] != str[nowPos] || visit[r][c] {
		return false
	}
	visit[r][c] = true
	for i := 0; i < 4; i++ {
		if isTharOk(matrix, r+direct[i][0], c+direct[i][1], visit, direct, str, nowPos+1) {
			return true
		}
	}
	visit[r][c] = false
	return false
}

func buildMatrix(array string, rows, cols int) [][]byte {
	strMatrix := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		strMatrix[i] = make([]byte, cols)
	}
	strByte := []byte(array)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			strMatrix[i][j] = strByte[rows*i+j]
		}
	}
	return strMatrix
}

func main() {
	arr := "abcdefg"
	rows := 1
	cols := len(arr)
	str := "abcd"
	fmt.Println(hasPath(arr, rows, cols, str))
}
