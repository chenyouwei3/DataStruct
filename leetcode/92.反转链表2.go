package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	var head = new(ListNode)
	head.Val = 0
	var node1 = new(ListNode)
	node1.Val = 1
	var node2 = new(ListNode)
	node2.Val = 2
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
	ll := reverseBetween(head, 2, 4)
	for ll != nil {
		fmt.Println("---->", ll.Val)
		ll = ll.Next
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
	}
	fmt.Println("找到了", pre.Val)
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newNode.Next
}
