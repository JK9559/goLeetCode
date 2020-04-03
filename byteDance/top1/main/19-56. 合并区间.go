package main

import "sort"

type Interval [][]int

func (in Interval) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in Interval) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}

func (in Interval) Len() int {
	return len(in)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return intervals
	}
	res := make([][]int, 0)
	sort.Sort(Interval(intervals))
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 || (len(res) >= 1 && res[len(res)-1][1] < intervals[i][0]) {
			res = append(res, intervals[i])
		} else {
			res[len(res)-1][1] = maxIn(res[len(res)-1][1], intervals[i][1])
		}
	}
	return res
}

func maxIn(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

}
