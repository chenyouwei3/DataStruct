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
	test := deleteDuplicatesA(head)
	for test != nil {
		fmt.Printf("%d->", test.Val)
		test = test.Next
	}
}

func deleteDuplicatesA(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	head = dummy
	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			//---------//
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
