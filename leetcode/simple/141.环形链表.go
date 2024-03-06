package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l1.Val = 3
	l2 := new(ListNode)
	l2.Val = 2
	l3 := new(ListNode)
	l3.Val = 0
	l4 := new(ListNode)
	l4.Val = -4
	l1.Next = l2
	l2.Next = l3
	l3.Next = l2
	fmt.Println(hasCycle(l1))
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return false
		}
	}
	return true
}
