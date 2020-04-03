package main

import (
	"container/list"
	"fmt"
)

type stack1 struct {
	list *list.List
}

func (s *stack1) push(val interface{}) {
	s.list.PushBack(val)
}

func (s *stack1) pop() interface{} {
	if s.list.Len() != 0 {
		e := s.list.Back()
		s.list.Remove(e)
		return e.Value
	}
	return nil
}

type myList struct {
	in  *stack1
	out *stack1
}

func newMyList() *myList {
	in := &stack1{list: list.New()}
	out := &stack1{list: list.New()}
	return &myList{in: in, out: out}
}

func (ml *myList) push(val int) {
	ml.in.push(val)
}

func (ml *myList) pop() int {
	if ml.out.list.Len() != 0 {
		return ml.out.pop().(int)
	} else {
		for ml.in.list.Len() != 0 {
			ml.out.push(ml.in.pop())
		}
		return ml.out.pop().(int)
	}
}

func main() {
	ml := newMyList()

	ml.push(1)
	ml.push(2)
	ml.push(3)
	fmt.Println(ml.pop())
	fmt.Println(ml.pop())
	fmt.Println(ml.pop())
}
