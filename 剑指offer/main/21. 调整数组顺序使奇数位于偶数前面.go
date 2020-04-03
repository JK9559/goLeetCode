package main

import "fmt"

func reOrderArray(nums *[]int) {
	l := len(*nums)
	for i := l - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if isEven((*nums)[j]) && !isEven((*nums)[j+1]) {
				(*nums)[j], (*nums)[j+1] = (*nums)[j+1], (*nums)[j]
			}
		}
	}
}

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	nums := []int{2, 1, 3, 4}

	reOrderArray(&nums)

	fmt.Println(nums)

}
