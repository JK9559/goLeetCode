package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	l, h := 0, len(letters)-1
	// 当l > h的时候 正好分布在target一右一左
	for l <= h {
		m := l + (h-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			h = m - 1
		}
	}
	fmt.Println(l, h)
	if l < len(letters) {
		return letters[l]
	} else {
		return letters[0]
	}
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g')))
}
