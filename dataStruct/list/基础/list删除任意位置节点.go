package main

import "fmt"

func main() {
	node1 := &ListNode{data: 1}
	node2 := &ListNode{data: 2}
	node3 := &ListNode{data: 3}
	node1.next = node2
	node2.next = node3
	//删除该节点的后面
	DeletedNode(node2)
	for node1 != nil {
		fmt.Printf("->%d", node1.data)
		node1 = node1.next
	}
}

func DeletedNode(nodeT *ListNode) {
	if nodeT.next == nil {
		return
	}
	Temp2 := nodeT.next.next
	nodeT.next = Temp2
}
