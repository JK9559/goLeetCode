package main

import "fmt"

/*
转换为树的问题
*/
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	help(len(nums), &nums, 0, &res)
	return res
}

func help(n int, nums *[]int, first int, res *[][]int) {
	if n == first {
		t := make([]int, len(*nums))
		copy(t, *nums)
		*res = append(*res, t)
	}
	for i := first; i < n; i++ {
		(*nums)[i], (*nums)[first] = (*nums)[first], (*nums)[i]
		help(n, nums, first+1, res)
		(*nums)[i], (*nums)[first] = (*nums)[first], (*nums)[i]
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
