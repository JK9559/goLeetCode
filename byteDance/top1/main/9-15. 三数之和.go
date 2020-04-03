/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	m := sort.IntSlice(nums)
	sort.Sort(m)
	nums = m
	for k := 0; k < len(nums); k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		target := 0 - nums[k]
		i, j := k+1, len(nums)-1
		for i < j {
			if nums[i]+nums[j] == target {
				res = append(res, [][]int{{nums[k], nums[i], nums[j]}}...)
				for i < j && nums[i] == nums[i+1] {
					i++
				}
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
