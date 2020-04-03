package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	n := len(lists)
	for n > 1 {
		k := (n + 1) / 2
		for i := 0; i < n/2; i++ {
			lists[i] = mergeTwoList(lists[i], lists[i+k])
		}
		n = k
	}
	return lists[0]
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
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
