package main

import "strconv"

/*
给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。
*/

func diffWaysToCompute(input string) []int {
	res := make([]int, 0)
	for k, v := range input {
		if v == '+' || v == '-' || v == '*' {
			lefts := diffWaysToCompute(input[0:k])
			rights := diffWaysToCompute(input[k+1:])
			for _, left := range lefts {
				for _, right := range rights {
					switch v {
					case '+':
						res = append(res, left+right)
						break
					case '-':
						res = append(res, left-right)
						break
					case '*':
						res = append(res, left*right)
						break
					}
				}
			}
		}
	}
	if len(res) == 0 {
		val, _ := strconv.Atoi(input)
		res = append(res, val)
	}
	return res
}

func main() {

}
