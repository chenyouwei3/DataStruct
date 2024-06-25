package main

import "fmt"

type ListNode struct {
	Next *ListNode
	Val  int
}

func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 2
	var node2 = new(ListNode)
	node2.Val = 3
	var node3 = new(ListNode)
	node3.Val = 4
	var node4 = new(ListNode)
	node4.Val = 5
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	test := reverseBetween(head, 2, 4)
	for test != nil {
		fmt.Printf("%d->", test.Val)
		test = test.Next
	}
}
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	newNode := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := newNode
	for i := 0; i < left-1 && pre.Next != nil; i++ {
		pre = pre.Next
		fmt.Println(pre)
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newNode.Next
}
