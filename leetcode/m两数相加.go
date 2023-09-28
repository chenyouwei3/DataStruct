package main

import "fmt"

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	//head := &ListNode{Val: 2}
	//node1 := &ListNode{Val: 4}
	//node2 := &ListNode{Val: 3}
	//
	//head.Next = node1
	//node1.Next = node2
	//
	//Head := &ListNode{Val: 5}
	//Node1 := &ListNode{Val: 6}
	//Node2 := &ListNode{Val: 4}
	////Node3 := &ListNode{Val: 1}
	//Head.Next = Node1
	//Node1.Next = Node2
	////Node2.Next = Node3
	head := &ListNode{Val: 9}
	node1 := &ListNode{Val: 9}
	node2 := &ListNode{Val: 9}
	node3 := &ListNode{Val: 9}
	node4 := &ListNode{Val: 9}
	node5 := &ListNode{Val: 9}
	node6 := &ListNode{Val: 9}
	node7 := &ListNode{Val: 9}
	//node8 := &ListNode{Val: 4}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	//node7.Next = node8

	Head := &ListNode{Val: 9}
	Node1 := &ListNode{Val: 9}
	Node2 := &ListNode{Val: 9}
	Node3 := &ListNode{Val: 9}
	Head.Next = Node1
	Node1.Next = Node2
	Node2.Next = Node3
	addTwoNumbers(head, Head)

}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	boll := false
	if l1.Val == 0 && l2.Val == 0 && l1.Next == nil && l2.Next == nil {
		l3 := &ListNode{Val: 0}
		return l3
	}
	cur1, cur2 := l1, l2
	l3 := new(ListNode)
	cur := l3
	for cur1 != nil && cur2 != nil {
		total := cur1.Val + cur2.Val
		if boll == true {
			total = total + 1
		}
		if total >= 10 {
			boll = true
			total = total % 10
		} else {
			boll = false
		}
		newnode := &ListNode{Val: total}
		cur1, cur2 = cur1.Next, cur2.Next
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newnode
		if cur2 == nil {
			p := &ListNode{Val: 0}
			cur2 = p
		}
		if cur1 == nil {
			p := &ListNode{Val: 0}
			cur1 = p
		}
		if cur1.Val == 0 && cur2.Val == 0 && boll == false {
			if cur1.Next == nil && cur2.Next == nil {
				break
			}
		}
	}
	for l3.Next != nil {
		fmt.Printf("->%d", l3.Next.Val)
		l3.Next = l3.Next.Next
	}
	return l3
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	boll := false
	if l1.Val == 0 && l2.Val == 0 && l1.Next == nil && l2.Next == nil {
		l3 := &ListNode{Val: 0}
		return l3
	}
	cur1, cur2 := l1, l2
	l3 := new(ListNode)
	cur := l3
	for cur1 != nil && cur2 != nil {
		total := cur1.Val + cur2.Val
		if boll == true {
			total = total + 1
		}
		if total >= 10 {
			boll = true
			total = total % 10
		} else {
			boll = false
		}
		newnode := &ListNode{Val: total}
		cur1, cur2 = cur1.Next, cur2.Next
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newnode

		if cur2 == nil {
			p := &ListNode{Val: 0}
			cur2 = p
		}
		if cur1 == nil {
			p := &ListNode{Val: 0}
			cur1 = p
		}
		if cur1.Val == 0 && cur2.Val == 0 && boll == false {
			if cur1.Next == nil && cur2.Next == nil {
				break
			}
		}
	}
	l4 := l3.Next
	return l4
}
