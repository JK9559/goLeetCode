package main

import (
	"container/list"
	"fmt"
)

/*
从尾到头反过来打印出每个结点的值。
*/

/*
使用栈
*/

type stack struct {
	list *list.List
}

func newStack() *stack {
	return &stack{list: list.New()}
}

func (s *stack) push(val interface{}) {
	s.list.PushBack(val)
}

func (s *stack) pop() interface{} {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

func (s *stack) isEmpty() bool {
	return s.list.Len() == 0
}

func printListFromTailToHead(node *ListNode) []int {
	st := newStack()
	ret := []int{}
	for node != nil {
		st.push(node.val)
		node = node.next
	}
	for !st.isEmpty() {
		ret = append(ret, st.pop().(int))
	}
	return ret
}

func main() {
	node := &ListNode{
		val: 1,
		next: &ListNode{
			val: 2,
			next: &ListNode{
				val:  3,
				next: nil,
			},
		},
	}

	fmt.Println(printListFromTailToHead(node))

}
