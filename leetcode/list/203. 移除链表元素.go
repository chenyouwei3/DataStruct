package main

import "fmt"

//
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
	l4.Val = 12
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	ll := removeElements(l1, 1)
	for ll != nil {
		fmt.Printf("->%d", ll.Val)
		ll = ll.Next
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	tempNode := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := tempNode
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return tempNode.Next
}
