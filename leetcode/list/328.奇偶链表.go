package main

import (
	"fmt"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	l1 := new(ListNode)
	l1.Val = 2
	l2 := new(ListNode)
	l2.Val = 1
	l3 := new(ListNode)
	l3.Val = 3
	l4 := new(ListNode)
	l4.Val = 5
	l5 := new(ListNode)
	l5.Val = 6
	l6 := new(ListNode)
	l6.Val = 4
	l7 := new(ListNode)
	l7.Val = 7

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	ll := oddEvenList(l1)
	for ll != nil {
		fmt.Println("---------->", ll.Val)
		ll = ll.Next
	}
}

func oddEvenList(head *ListNode) *ListNode {
	var left, right, res []int
	pre := new(ListNode)
	cur := pre
	for i := 0; head != nil; i++ {
		fmt.Println(i, head.Val)
		if i%2 == 0 {
			left = append(left, head.Val)
		} else {
			right = append(right, head.Val)
		}
		head = head.Next
	}
	res = append(left, right...)
	for _, v := range res {
		pre.Next = &ListNode{Val: v}
		pre = pre.Next
	}
	return cur.Next
}
