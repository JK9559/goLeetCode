package main

// 操作给定的二叉树，将其变换为源二叉树的镜像。
func Mirror(root *TreeNode) {
	if root == nil {
		return
	}
	swap(root)
	Mirror(root.Left)
	Mirror(root.Right)
}

func swap(root *TreeNode) {
	t := root.Left
	root.Left = root.Right
	root.Right = t
}

func main() {

}
