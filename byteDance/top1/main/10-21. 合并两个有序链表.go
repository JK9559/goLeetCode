package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := pHead
	for nil != l1 && nil != l2 {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if nil == l1 {
		p.Next = l2
	} else {
		p.Next = l1
	}
	return pHead.Next
}
