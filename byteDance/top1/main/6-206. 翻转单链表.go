/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

type ListNode_23 struct {
	Val  int
	Next *ListNode_23
}

func reverseList(head *ListNode_23) *ListNode_23 {
	// 注意保持记录prev 和 next
	if head == nil {
		return head
	}
	var prev *ListNode_23
	var next *ListNode_23
	for head != nil {
		next = head.Next

		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {
	var node1 = ListNode_23{
		Val: 1,
		Next: &ListNode_23{
			Val: 2,
			Next: &ListNode_23{
				Val: 3,
				Next: &ListNode_23{
					Val: 4,
					Next: &ListNode_23{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	var res = reverseList(&node1)
	for {
		fmt.Println(res.Val)
		if res.Next == nil {
			break
		}
		res = res.Next
	}
}
