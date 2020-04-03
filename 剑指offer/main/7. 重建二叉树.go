package main

import (
	"fmt"
	"time"
)

/*
根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*/

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func buildTree1(preorder, inorder []int) *TreeNode1 {
	indexMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		indexMap[inorder[i]] = i
	}
	return helper(preorder, 0, len(preorder)-1, 0, &indexMap)
}

func helper(pre []int, preLeft, preRight, inLeft int, indexMap *map[int]int) *TreeNode1 {
	if preLeft > preRight {
		return nil
	}
	//fmt.Println(preLeft)
	inIndex := (*indexMap)[pre[preLeft]]
	leftLen := inIndex - inLeft
	node := &TreeNode1{
		Val:   pre[preLeft],
		Left:  helper(pre, preLeft+1, preLeft+leftLen, inLeft, indexMap),
		Right: helper(pre, preLeft+leftLen+1, preRight, inLeft+leftLen+1, indexMap),
	}
	return node
}

func buildTree(preorder, inorder []int) *TreeNode1 {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	var i int
	for i = range inorder {
		if preorder[0] == inorder[i] {
			break
		}
	}
	node := &TreeNode1{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+i], inorder[:i]),
		Right: buildTree(preorder[i+1:], inorder[i+1:]),
	}
	return node
}

func main() {
	intMax := 1 << 10
	pre := make([]int, intMax)
	in := make([]int, intMax)
	for i := 0; i < intMax; i++ {
		in[i] = i
		pre[i] = intMax - i
	}

	now := time.Now()
	buildTree1(pre, in)
	fmt.Println(time.Since(now))

	now = time.Now()
	buildTree(pre, in)
	fmt.Println(time.Since(now))

}
