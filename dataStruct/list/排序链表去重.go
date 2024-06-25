package main

import "fmt"

//type ListNode struct {
//	Next *ListNode
//	Val  int
//}

func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 2
	var node2 = new(ListNode)
	node2.Val = 2
	var node3 = new(ListNode)
	node3.Val = 2
	var node4 = new(ListNode)
	node4.Val = 3
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	test := deleteDuplicates(head)
	for test != nil {
		fmt.Printf("%d->", test.Val)
		test = test.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		for curr.Next != nil && curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
	return head
}
