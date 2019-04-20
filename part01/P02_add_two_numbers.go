package part01

import (
	"goLeetCode/datastructure"
)

func addTwoNumbers(l1 *datastructure.ListNode, l2 *datastructure.ListNode) *datastructure.ListNode {
	rst, trick := 0, 0
	p, q, head := l1, l2, new(datastructure.ListNode)
	pRst := head
	for ; ; {
		val1, val2, node := 0, 0, new(datastructure.ListNode)
		rst = 0
		if nil != p {
			val1 = p.Val
			p = p.Next
		}
		if nil != q {
			val2 = q.Val
			q = q.Next
		}
		rst = val1 + val2 + trick
		trick = rst / 10
		rst = rst % 10
		pRst.Val = rst
		if nil == p && nil == q && 0 == trick {
			pRst.Next = nil
			return head
		}
		pRst.Next = node
		pRst = node
	}
}

func P02() {
	l1 := datastructure.CreateListNode([]int{1})
	l2 := datastructure.CreateListNode([]int{9, 9, 9, 9, 9})
	datastructure.PrintListNode(addTwoNumbers(l1, l2))
}
