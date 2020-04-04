package main

/*
如果删除节点不是尾节点 则可以将用要删除节点的next覆盖当前节点
如果删除节点是尾节点：
 1. 如果只有一个节点 则将头节点删除
 2. 如果删除节点为尾节点 则必须遍历到要删除节点的前一个节点 将这个节点的next置为nil
*/

type ListNodeDel struct {
	val  int
	next *ListNodeDel
}

func deleteNode(head, delNode *ListNodeDel) *ListNodeDel {
	if head == nil || delNode == nil {
		return head
	}
	if delNode.next != nil {
		// 注意这波操作 删除节点 本应该会找到前置节点 将前置节点的next赋值为删除节点的next
		// 但是这里，将待删除节点的next的值 变更为了 待删除的next的next，相当于删除了待删除的节点的下一个节点
		tmpNode := delNode.next
		delNode.val = tmpNode.val
		delNode.next = tmpNode.next
	} else {
		if delNode == head {
			head = nil
			delNode = nil
		} else {
			node := head
			for node.next != delNode {
				node = node.next
			}
			node.next = nil
		}
	}
	return head
}

func main() {

}
