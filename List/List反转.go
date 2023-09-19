package main

import "fmt"

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
	//show3333(head)
	reverse(head)

}

func reverse(head *ListNode) {
	var cur, pre *ListNode
	cur = head
	for cur != nil {
		Temp := cur.next
		cur.next = pre
		pre = cur
		cur = Temp
	}
	for pre != nil {
		fmt.Printf("->%d", pre.data)
		pre = pre.next
	}
}
