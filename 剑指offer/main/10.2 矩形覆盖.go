package main

import "fmt"

/*
我们可以用 2*1 的小矩形横着或者竖着去覆盖更大的矩形。请问用 n 个 2*1 的小矩形无重叠地覆盖一个 2*n 的大矩形，总共有多少种方法？
*/

/*
当n为1时 一种方法
当n为2时 二种方法
当为n时，问题有两种形式：
1. 第一行横着摆一个2*1的矩形 剩下求n-1个2*1有多少种方式即可
2. 前两行为竖着摆二个1*2的矩形 剩下求n-2个2*1有多少种方式即可
所以 递推公式为：
f(n) = 1 n = 1
f(n) = 2 n = 2
f(n) = f(n-1) + f(n-2) n > 2
可知 为fibonacci数列
*/

func RectCover(n int) int {
	if n < 0 {
		return 0
	} else if n <= 2 {
		return n
	} else {
		a, b := 1, 2
		ret := 0
		for i := 2; i < n; i++ {
			ret = a + b
			a = b
			b = ret
		}
		return ret
	}
}

func main() {
	fmt.Println(RectCover(3))
}
