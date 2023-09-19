package main

import "fmt"

//type ListNode struct {
//	data int
//	next *ListNode
//}

func main() {
	var head *ListNode
	var Head *ListNode

	// 初始化链表
	head = &ListNode{data: 2}
	node1 := &ListNode{data: 5}
	node2 := &ListNode{data: 3}
	head.next = node1
	node1.next = node2

	// 初始化链表
	Head = &ListNode{data: 5}
	Node1 := &ListNode{data: 6}
	Node2 := &ListNode{data: 4}
	Head.next = Node1
	Node1.next = Node2
	showNode1(head)

}

// 遍历列表
func showNode1(p *ListNode) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
