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
	l2.Val = 4
	l3 := new(ListNode)
	l3.Val = 3
	l4 := new(ListNode)
	l4.Val = 2
	l5 := new(ListNode)
	l5.Val = 5
	l6 := new(ListNode)
	l6.Val = 2

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	partitionTest(l1, 3)

}

func partitionTest(head *ListNode, x int) {
	left := &ListNode{Val: 0}
	right := &ListNode{Val: 0}
	before := left
	after := right
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = right.Next
	for left != nil {
		fmt.Println("*->", left.Val)
		left = left.Next
	}
	for left != nil {
		fmt.Println("*->", left.Val)
		left = left.Next
	}
	return
}

func partition(head *ListNode, x int) *ListNode {
	left := &ListNode{Val: 0}
	right := &ListNode{Val: 0}
	before := left
	after := right
	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = right.Next
	return left.Next
}
