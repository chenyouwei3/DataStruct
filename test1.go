package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 2}
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 3}
	head.Next = node1
	node1.Next = node2
	Head := &ListNode{Val: 5}
	Node1 := &ListNode{Val: 6}
	Node2 := &ListNode{Val: 4}
	Head.Next = Node1
	Node1.Next = Node2
	l := head
	l1 := l
	l.Val = 1
	for l1 != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
	//for l != nil {
	//	fmt.Println(l.Val)
	//	l = l.Next
	//}
}
