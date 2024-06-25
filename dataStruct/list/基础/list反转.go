package main

import "fmt"

//
//type ListNode struct {
//	data int
//	next *ListNode
//}

func main() {
	var head = new(ListNode)
	head.data = 1
	var node1 = new(ListNode)
	node1.data = 2
	var node2 = new(ListNode)
	node2.data = 3

	head.next = node1

	node1.next = node2
	res := reverse(head)
	for res != nil {
		fmt.Printf("%d->", res.data)
		res = res.next
	}
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.next //中间变量
		head.next = pre   //移动方向
		pre = head        //更新节点
		head = temp       //更新节点
	}
	return pre
}

func reverseTest(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.next
		head.next = pre
		pre = head
		head = temp
	}
	return pre
}
