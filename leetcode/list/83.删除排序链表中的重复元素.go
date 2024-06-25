package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 1
	l3 := new(ListNode)
	l3.Val = 1
	l4 := new(ListNode)
	l4.Val = 2
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	ll := deleteDuplicates2(l1)
	for ll != nil {
		fmt.Printf("->%d", ll.Val)
		ll = ll.Next
	}
}

func deleteDuplicates2(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}
