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
	ll := deleteDuplicates(l1)
	for ll != nil {
		fmt.Printf("->%d", ll.Val)
		ll = ll.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
