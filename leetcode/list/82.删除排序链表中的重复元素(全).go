package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 3
	var node2 = new(ListNode)
	node2.Val = 4
	var node3 = new(ListNode)
	node3.Val = 4
	var node4 = new(ListNode)
	node4.Val = 5
	//var node5 = new(ListNode)
	//node5.Val = 6
	//var node6 = new(ListNode)
	//node6.Val = 7
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node5
	//node5.Next = node6
	ll := deleteDuplicates1(head)
	for ll != nil {
		fmt.Println("---->", ll.Val)
		ll = ll.Next
	}

}

func deleteDuplicates1(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	head = dummy
	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}
