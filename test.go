package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func main() {
	head := &ListNode{data: 1}
	node1 := &ListNode{data: 2}
	node2 := &ListNode{data: 3}
	head.next = node1
	node1.next = node2
	test(5, head, 1)
}

func test(data int, head *ListNode, index int) {
	newNode := &ListNode{
		data: data,
	}
	cur := head
	if index == 0 {
		newNode.next = head
		for newNode != nil {
			fmt.Printf("%d->", newNode.data)
			newNode = newNode.next
		}
	}
	var pre *ListNode
	cout := 0
	for cur != nil && cout < index {
		Temp := cur.next
		pre = cur
		cur = Temp
		cout++
	}
	pre.next = newNode
	newNode.next = cur

	for head != nil {
		fmt.Printf("-->%d", head.data)
		head = head.next
	}
}
