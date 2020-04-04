package main

import "fmt"

type ListNodeDelII struct {
	val  int
	next *ListNodeDelII
}

func deleteDuplication(head *ListNodeDelII) *ListNodeDelII {
	if head == nil || head.next == nil {
		return head
	}
	nextNode := head.next
	if head.val == nextNode.val {
		// 如果当前头结点 和 下一个节点val相等 那么就一直向后找 直到返回不相等的那个节点
		for nextNode != nil && nextNode.val == head.val {
			nextNode = nextNode.next
		}
		return deleteDuplication(nextNode)
	} else {
		// 否则 当前节点的next为以下一个节点为首的结果
		// 为了不丢失这个不同的头结点
		head.next = deleteDuplication(head.next)
		return head
	}
}

func main() {
	node := &ListNodeDelII{
		val: 2,
		next: &ListNodeDelII{
			val: 2,
			next: &ListNodeDelII{
				val:  3,
				next: nil,
			},
		},
	}

	rstNode := deleteDuplication(node)

	for rstNode != nil {
		fmt.Println(rstNode.val)
		rstNode = rstNode.next
	}

}
