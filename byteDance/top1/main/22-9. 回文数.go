package main

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
*/

// 翻转数字 不使用字符串 只翻转一半的数字 判断如果原始数字小于了翻转的数字 那么就说明已经处理了一半的数字
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	resInt := 0
	for x > resInt {
		resInt = resInt*10 + x%10
		x /= 10
	}
	return x == resInt || x == resInt/10
}

func main() {

}
