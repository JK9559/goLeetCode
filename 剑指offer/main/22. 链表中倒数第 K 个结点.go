package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func FindKthToTail(head *ListNode, k int) *ListNode {
	if nil == head {
		return nil
	}
	p1, p2 := head, head
	for p1 != nil && k > 0 {
		p1 = p1.next
		k--
	}
	if k < 0 {
		return nil
	}
	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2
}

func main() {
	list := &ListNode{
		val: 1,
		next: &ListNode{
			val: 2,
			next: &ListNode{
				val: 3,
				next: &ListNode{
					val:  4,
					next: nil,
				},
			},
		},
	}

	fmt.Println(FindKthToTail(list, 2).val)
}
