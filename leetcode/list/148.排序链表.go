package main

import (
	"fmt"
	"sort"
)

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l3 := new(ListNode)
	l3.Val = 4
	l4 := new(ListNode)
	l4.Val = 3
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	ll := sortList(l1)
	for ll != nil {
		fmt.Printf("->%d", ll.Val)
		ll = ll.Next
	}
}

func sortList(head *ListNode) *ListNode {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	sort.Ints(res)
	cur := new(ListNode)
	pre := &ListNode{
		Val:  0,
		Next: cur,
	}
	for _, v := range res {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return pre.Next.Next
}
