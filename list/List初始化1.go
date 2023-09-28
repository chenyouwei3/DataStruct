package main

import "fmt"

//type ListNode struct {
//	data int
//	next *ListNode
//}

// 遍历列表
func showNode(p *ListNode) {
	for p != nil {
		fmt.Printf("%d->", p.data)
		p = p.next
	}
}

func main() {
	var head = new(ListNode)
	head.data = 1
	var node1 = new(ListNode)
	node1.data = 2
	var node2 = new(ListNode)
	node2.data = 3

	head.next = node1
	node1.next = node2
	fmt.Println(head.next)
	showNode(head)
}
