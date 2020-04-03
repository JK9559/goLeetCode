package main

// 输入两棵二叉树A，B，判断B是不是A的子结构。（ps：我们约定空树不是任意一个树的子结构）

func HasSubtree(root1, root2 *TreeNode) bool {
	if nil == root1 || nil == root2 {
		return false
	}
	return isSubtreeWithRoot(root1, root2) || HasSubtree(root1.Left, root2) || HasSubtree(root1.Right, root2)
}

// 判断两棵树是不是完全相同
func isSubtreeWithRoot(root1, root2 *TreeNode) bool {
	if nil == root1 {
		return false
	}
	if nil == root2 {
		return true
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSubtreeWithRoot(root1.Left, root2.Left) && isSubtreeWithRoot(root1.Right, root2.Right)
}

func main() {

}
