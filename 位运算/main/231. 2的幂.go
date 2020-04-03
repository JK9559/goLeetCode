package main

/*
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
*/

/*
2的幂只有一个1
*/

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}

func main() {

}
