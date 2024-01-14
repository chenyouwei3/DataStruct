package main

func swapPairs(head *ListNode) *ListNode {
	newNode := &ListNode{Next: head}
	node0 := newNode
	node1 := head
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next //2
		node3 := node2.Next //3

		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		//交换节点
		node0 = node1
		node1 = node3
	}

	return newNode.Next
}
