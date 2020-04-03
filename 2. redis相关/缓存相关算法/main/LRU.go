/*
 * Copyright © 2019 wkang. All rights reserved.
 */

package main

import "fmt"

/*
LRU是啥？最近最少使用 将cache头作为最常访问点 当有节点被访问时 将此节点设置为最常访问节点（置于头部）
当位置不够时 删除最近最少访问的节点（也就是最尾部节点）
*/

// 双向链表的节点
type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

// 双向链表
type DLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDLinkedList(head *Node, tail *Node) *DLinkedList {
	return &DLinkedList{Head: head, Tail: tail}
}

// 当链表的头结点==尾节点==nil的时候 链表为空
func (d *DLinkedList) isEmpty() bool {
	if d.Head == nil && d.Tail == nil {
		return true
	} else {
		return false
	}
}

func (d *DLinkedList) remove(node *Node) {
	// 当链表里只有一个节点时
	if d.Head == d.Tail {
		d.Head = nil
		d.Tail = nil
		return
	}
	// 当移除的节点是头结点的时候
	if node == d.Head {
		d.Head = node.Next
		node.Next.Pre = nil
		return
	}
	// 当移除的节点是尾节点的时候
	if node == d.Tail {
		d.Tail = node.Pre
		node.Pre.Next = nil
		return
	}
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (d *DLinkedList) removeTail() {
	if d.Tail != nil {
		d.remove(d.Tail)
	}
}

func (d *DLinkedList) addHead(node *Node) {
	// 添加头结点
	if d.Head == nil {
		d.Head = node
		d.Tail = node
		node.Pre = nil
		node.Next = nil
		return
	}
	node.Pre = nil
	d.Head.Pre = node
	node.Next = d.Head
	d.Head = node
}

type LRUCache struct {
	cap   int
	size  int
	Cache *DLinkedList
	HsMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:   capacity,
		size:  0,
		Cache: NewDLinkedList(nil, nil),
		HsMap: make(map[int]*Node),
	}
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.HsMap[key]; ok {
		this.Cache.remove(val)
		this.Cache.addHead(val)
		return val.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.HsMap[key]; ok {
		this.Cache.remove(val)
		val.Val = value
		this.Cache.addHead(val)
	} else {
		node := &Node{
			Key:  key,
			Val:  value,
			Pre:  nil,
			Next: nil,
		}
		this.HsMap[key] = node
		this.Cache.addHead(node)
		this.size += 1
		if this.size > this.cap {
			this.size -= 1
			delete(this.HsMap, this.Cache.Tail.Key)
			this.Cache.removeTail()
		}
	}
}

func main() {
	capacity := 2
	cache := Constructor(capacity)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4
}
