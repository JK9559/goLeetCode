package datastructure

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(array []int) *ListNode {
	head := new(ListNode)
	head.Val = array[0]
	p := head
	for i := 1; i < len(array); i++ {
		node := new(ListNode)
		p.Next = node
		node.Val = array[i]
		p = node
	}
	p.Next = nil
	return head
}

func PrintListNode(lNode *ListNode) {

	for p := lNode; nil != p; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}

func TestListNode() {
	array := []int{2, 4, 3}
	listNode := CreateListNode(array)
	PrintListNode(listNode)
}
