package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	key1 := (len1 + len2 + 1) / 2
	key2 := (len1 + len2 + 2) / 2
	val1 := int(^uint(0) >> 1)
	val2 := int(^uint(0) >> 1)
	fmt.Println(val1)
	i, j, len := 0, 0, 1
	for i < len1 && j < len2 {
		if nums1[i] < nums2[j] {
			if len == key1 {
				val1 = nums1[i]
			}
			if len == key2 {
				val2 = nums1[i]
			}
			i++
		} else {
			if len == key1 {
				val1 = nums2[j]
			}
			if len == key2 {
				val2 = nums2[j]
			}
			j++
		}
		len++
	}
	if len <= key1 {
		// 取两个值
		// 当前大数组里的数据长度为 len1+j, 到达第key1个元素的偏移量为key1-(len1+j),
		// 对于未遍历完成的数组已经即将处理到元素nums[j](也就是说nums[j]不包含在len1+j个元素中)
		// 又 下标从0开始计数 个数从1开始计数 所以个数转下标时 减去1所以 元素下标为 j + (key1-(len1+j)) -1
		if i >= len1 {
			val1 = nums2[key1-len1-1]
			val2 = nums2[key2-len1-1]
		}
		if j >= len2 {
			val1 = nums1[key1-len2-1]
			val2 = nums1[key2-len2-1]
		}
	} else if len == key2 {
		// 取一个值
		if i >= len1 {
			val2 = nums2[j]
		}
		if j >= len2 {
			val2 = nums1[i]
		}
	}
	return float64(val1+val2) / 2
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func find(nums1 []int, nums2 []int) float64 {
	// 题解见 https://blog.csdn.net/aikudexue/article/details/90729786
	m := len(nums1)
	n := len(nums2)
	if m > n {
		nums2, nums1, m, n = nums1, nums2, n, m
	}
	if m+n == 1 {
		if m == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	mini, maxi, num := 0, m, (m+n+1)/2

	for mini <= maxi {
		i := (mini + maxi) / 2
		j := num - i
		if i > 0 && nums1[i-1] > nums2[j] {
			maxi = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			mini = i + 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			var minRight int
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			} else {
				return float64(maxLeft+minRight) / 2.0
			}
		}
	}
	return 0
}

func P04() {
	ary1 := []int{}
	ary2 := []int{1, 2}
	rst := findMedianSortedArrays(ary1, ary2)
	fmt.Println(rst)
	rs1 := find(ary1, ary2)
	fmt.Println(rs1)
}
