package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 3
	var node2 = new(ListNode)
	node2.Val = 4
	var node3 = new(ListNode)
	node3.Val = 4
	var node4 = new(ListNode)
	node4.Val = 5
	//var node5 = new(ListNode)
	//node5.Val = 6
	//var node6 = new(ListNode)
	//node6.Val = 7
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node5
	//node5.Next = node6
	ll := deleteDuplicates(head)
	for ll != nil {
		fmt.Println("---->", ll.Val)
		ll = ll.Next
	}

}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	cur := head
	temp := head.Val
	cur = cur.Next
	for cur != nil {
		if cur.Val == temp {
			pre.Next = cur.Next
		} else {
			pre = cur
			temp = cur.Val
		}
		cur = cur.Next
	}
	return newHead.Next
}
