package main

import "fmt"

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		// 移动一位
		res <<= 1
		// 设置值
		res |= num & 1
		// 取num的下一位
		num >>= 1
	}
	return res
}

func reverseBits1(num uint32) uint32 {
	res := uint32(0)
	for i := 0; i < 32; i++ {
		res |= num & 1
		fmt.Printf("%b\n", res)
		res <<= 1
		fmt.Printf("%b\n", res)
		num >>= 1
	}
	return res
}

func main() {
	fmt.Printf("%b\n", reverseBits(5))
	fmt.Printf("%b\n", reverseBits1(5))
}
