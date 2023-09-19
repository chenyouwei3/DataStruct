package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func main() {
	var head = new(ListNode)
	head.data = 1
	var node1 = new(ListNode)
	node1.data = 2
	var node2 = new(ListNode)
	node2.data = 3

	head.next = node1
	node1.next = node2
	InsertAtIndexList(head, 8, 3)
	//for head != nil {
	//	fmt.Printf("%d->", head.data)
	//	head = head.next
	//}
	//reverse(head)
}

func InsertAtIndexList(head *ListNode, data int, index int) {
	newNode := &ListNode{
		data: data,
		next: nil,
	}
	cur := head
	if index == 0 {
		newNode.next = head
		for newNode != nil {
			fmt.Printf("-->%d", newNode.data)
			newNode = newNode.next
		}
	}
	var pre *ListNode
	cout := 0
	for cur != nil && cout < index {
		//往后面遍历
		Temp := cur.next
		pre = cur
		cur = Temp
		//
		cout++
		fmt.Println(pre, "pre") //1
		fmt.Println(cur, "cur") //2
	}
	pre.next = newNode
	newNode.next = cur

	for head != nil {
		fmt.Printf("-->%d", head.data)
		head = head.next
	}

}
