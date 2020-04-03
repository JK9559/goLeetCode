package main

import (
	"fmt"
	"strings"
)

/*
给定两个字符串 s1 和 s2，要求判定 s2 是否能够被 s1 做循环移位得到的字符串包含。
*/
/*
s1 = AABCD, s2 = CDAA
Return : true
*/
func isContain(s1, s2 string) bool {
	return !(-1 == strings.LastIndex(s1+s1, s2))
}

func main() {
	fmt.Println(isContain("322", "223"))
}
