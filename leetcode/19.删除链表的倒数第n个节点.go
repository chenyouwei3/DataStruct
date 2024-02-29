package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := reverse0(head)
	Temp := cur
	if n == 1 {
		res0 := Temp.Next
		return reverse0(res0)
	}
	for i := 2; i < n; i++ {
		Temp = Temp.Next
	}
	Temp.Next = Temp.Next.Next
	res := reverse0(cur)
	return res
}

func reverse0(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		Temp := head.Next
		head.Next = pre
		pre = head
		head = Temp
	}
	return pre
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
	var node5 = new(ListNode)
	node5.Val = 6
	var node6 = new(ListNode)
	node6.Val = 7
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	res := removeNthFromEnd(head, 1)
	for res != nil {
		fmt.Printf("->%d", res.Val)
		res = res.Next
	}
}
