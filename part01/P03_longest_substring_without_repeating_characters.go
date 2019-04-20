package part01

import "fmt"

// 滑动窗口,设置数组chs记录每个字符串最后出现的位置
// left为最左端 i为当前位置
// 滑动left,规则为当右侧遇到字符已经包含在left-i时(判断规则为，查找chs数组，如果上一个该字符的位置大于等于left就是包含了)
// 那么移动left到字符的下一位
// 需要针对空字符串做特殊处理
func lengthOfLongestSubstring(s string) int {
	if 0 == len(s) {
		return 0
	}
	chs := make([]int, 256)
	for i := 0; i < len(chs); i++ {
		chs[i] = -1
	}
	left, res := 0, 1
	for i := 0; i < len(s); i++ {
		lastPos := chs[s[i]]
		chs[s[i]] = i
		if lastPos >= left {
			left = lastPos + 1
		}
		res = maxInt(res, i-left+1)
	}
	return res
}

func maxInt(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func P03() {
	fmt.Println(lengthOfLongestSubstring("abc"))
}
