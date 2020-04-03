package main

import "fmt"

// 4的幂 只有奇数位有1
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && (num&0x55555555) != 0
}

func main() {
	fmt.Println(isPowerOfFour(4))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(2))
	fmt.Println(16 & 0x55555555)
}
