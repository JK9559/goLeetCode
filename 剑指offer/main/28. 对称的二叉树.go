package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSymmetrical(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return isSymmetricalSub(pRoot.Left, pRoot.Right)
}

func isSymmetricalSub(p1, p2 *TreeNode) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
		return false
	}
	if p1.Val != p2.Val {
		return false
	}
	return isSymmetricalSub(p1.Left, p1.Right) && isSymmetricalSub(p2.Left, p2.Right)
}

func main() {
	pr := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(isSymmetrical(pr))
}
