package main

/*
模拟竖式加法
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var trick = 0
	var head, tail *ListNode
	head = &ListNode{}
	tail = head
	for {
		if l1 == nil && l2 == nil && trick == 0 {
			break
		}
		// 因为go没有三元操作符 所以为一个值赋默认值为0
		var v1 int
		if v1 = 0; l1 != nil {
			v1 = l1.Val
		}
		var v2 int
		if v2 = 0; l2 != nil {
			v2 = l2.Val
		}
		val := v1 + v2 + trick
		trick = val / 10
		tail.Val = val % 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		// 只有当满足条件 才会有新的节点
		if l1 != nil || l2 != nil || trick != 0 {
			tail.Next = &ListNode{}
			tail = tail.Next
		}
	}
	return head
}

func main() {

}
