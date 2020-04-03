package main

func hasAlternatingBits(n int) bool {
	dif := n ^ (n >> 1)
	return dif&(dif+1) == 0
}

func main() {

}
