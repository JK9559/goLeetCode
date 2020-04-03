package main

import "fmt"

/*
一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
*/

/*
1. 当青蛙第一次跳1级台阶 剩下n-1级 有 f(n-1) 种
2. 当青蛙第一次跳2级台阶 剩下n-2级 有 f(n-2) 种
所以 青蛙跳n级台阶 有 f(n-1) + f(n-2) 种
*/

func JumpFloor(n int) int {
	if n < 0 {
		return 0
	} else if n <= 2 {
		return n
	} else {
		ret, a, b := 0, 1, 2
		for i := 2; i < 3; i++ {
			ret = a + b
			a = b
			b = ret
		}
		return ret
	}
}

func main() {
	fmt.Println(JumpFloor(3))
}
