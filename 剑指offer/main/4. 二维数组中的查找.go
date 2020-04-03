package main

import "fmt"

/*
给定一个二维数组，其每一行从左到右递增排序，从上到下也是递增排序。给定一个数，判断这个数是否在该二维数组中。
*/

/*
要求时间复杂度 O(M + N)，空间复杂度 O(1)。其中 M 为行数，N 为 列数。

该二维数组中的一个数，小于它的数一定在其左边，大于它的数一定在其下边。
因此，从右上角开始查找，就可以根据 target 和当前元素的大小关系来缩小查找区间，当前元素的查找区间为左下角的所有元素。
*/

/*
为啥选择右上角 有啥好处？因为，比这个值大的只能往下，比这值小的，只能往左
*/

func find(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var i, j = 0, len(matrix[0]) - 1
	for i <= len(matrix)-1 && j >= 0 {
		if target < matrix[i][j] {
			j--
		} else if target > matrix[i][j] {
			i++
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}

	fmt.Println(find(matrix, 5))
	fmt.Println(find(matrix, 20))
}
