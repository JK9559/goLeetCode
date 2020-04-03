package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	pre, end := dummy, dummy
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		// start 为要反转的首节点 next 为下次要反转的首节点 end 为要反转的终点 pre 为要反转首节点的前一个节点
		start, next := pre.Next, end.Next
		end.Next = nil
		// 翻转start为首的链表
		pre.Next = reverse(start)
		// 这时start为反转后链表的尾节点 与next相连
		start.Next = next
		// 复位 pre 和 end 为下一次要反转链表的前一个节点
		pre, end = start, start
	}
	return dummy.Next
}

// 翻转链表 单链表
func reverse(head *ListNode) *ListNode {
	pre, cur := &ListNode{
		Val:  -1,
		Next: nil,
	}, head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

func main() {

}
