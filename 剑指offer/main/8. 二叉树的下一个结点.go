package main

/*
给定一个二叉树和其中的一个结点，请找出中序遍历顺序的下一个结点并且返回 。注意，树中的结点不仅包含左右子结点，同时包含指向父结点的指针。
*/

/*
中序遍历的下一个节点是：
1. 如果一个节点的右子树不为空，那么该节点的下一个节点是右子树的最左节点
2. 否则，向上找第一个左链接指向的树包含该节点的祖先节点。
*/

type TreeLinkNode struct {
	val   int
	left  *TreeLinkNode
	right *TreeLinkNode
	// next 指向节点的父节点
	next *TreeLinkNode
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode.right != nil {
		p := pNode.right
		for p.left != nil {
			p = p.left
		}
		return p
	}
	for pNode.next != nil {
		parent := pNode.next
		if parent.left == pNode {
			return parent
		}
		pNode = pNode.next
	}
	return nil
}

func main() {

}
