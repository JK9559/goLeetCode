package main

/*
一个数 x 的开方 sqrt 一定在 0 ~ x 之间，并且满足 sqrt == x / sqrt。可以利用二分查找在 0 ~ x 之间查找 sqrt。
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	l, h := 1, x
	for l <= h {
		m := l + (h-l)/2
		sqrt := x / m
		if sqrt == m {
			return m
		} else if m < sqrt {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	return h
}

func main() {

}
