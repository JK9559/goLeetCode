package main

import "math"

/*
从ans * 10 + pop > MAX_VALUE这个溢出条件来看

    当出现 ans > MAX_VALUE / 10 且 还有pop需要添加 时，则一定溢出
    当出现 ans == MAX_VALUE / 10 且 pop > 7 时，则一定溢出，7是2^31 - 1的个位数

从ans * 10 + pop < MIN_VALUE这个溢出条件来看

    当出现 ans < MIN_VALUE / 10 且 还有pop需要添加 时，则一定溢出
    当出现 ans == MIN_VALUE / 10 且 pop < -8 时，则一定溢出，8是-2^31的个位数

*/

const MAX_INT = math.MaxInt32
const MIN_INT = math.MinInt32

func reverseInt(x int) int {
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if res > MAX_INT/10 || (res == MAX_INT/10 && pop > 7) {
			return 0
		}
		if res < MIN_INT/10 || (res == MIN_INT/10 && pop < -8) {
			return 0
		}
		res = res*10 + pop
	}
	return res
}
