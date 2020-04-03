package main

/*
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。
*/

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}
	return generateSubtrees(1, n)
}

func generateSubtrees(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		leftSubs := generateSubtrees(start, i-1)
		rightSubs := generateSubtrees(i+1, end)
		for _, leftSub := range leftSubs {
			for _, rightSub := range rightSubs {
				root := &TreeNode{
					Val:   i,
					Left:  leftSub,
					Right: rightSub,
				}
				res = append(res, root)
			}
		}
	}
	return res
}

func main() {

}
