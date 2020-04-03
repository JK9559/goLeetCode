package main

import "math"

/*
一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级... 它也可以跳上 n 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法。
*/

/*
f(n) = f(n-1) + f(n-2) +...+ f(0)
f(n-1) = f(n-2) +...+ f(0)
f(n) - f(n-1) = f(n-1)
f(n) = 2 * f(n-1)
f(1) = 1
*/

func JumpFloorII(n int) int {
	return int(math.Pow(2.0, float64(n-1)))
}

func main() {

}
