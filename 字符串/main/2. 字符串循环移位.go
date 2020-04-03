package main

import "fmt"

/*
将字符串向右循环移动 k 位。

将 abcd123 中的 abcd 和 123 单独翻转，得到 dcba321，然后对整个字符串进行翻转，得到 123abcd。
*/

/*
s = "abcd123" k = 3
Return "123abcd"
*/

func reversek(s string, k int) string {
	return rStr(rStr(s[:len(s)-k]) + rStr(s[len(s)-k:]))
}

// 反转字符串
func rStr(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reversek("abcd123", 3))
}
