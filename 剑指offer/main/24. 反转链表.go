package main

func ReverseList(head *ListNode) *ListNode {
	newList := &ListNode{-1, nil}
	for head != nil {
		p := head.next
		head.next = newList.next
		newList.next = head
		head = p
	}
	return newList.next
}

func main() {

}
