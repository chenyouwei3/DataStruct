package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l3 := new(ListNode)
	l3.Val = 4
	l4 := new(ListNode)
	l4.Val = 3
	l5 := new(ListNode)
	l5.Val = 5

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	deleteNode(l3)
	for l1 != nil {
		fmt.Println("---------->", l1.Val)
		l1 = l1.Next
	}
}

func deleteNode(node *ListNode) {
	pre := node
	var temp []int
	for pre != nil {
		temp = append(temp, pre.Val)
		pre = pre.Next
	}
	for i, v := range temp[1:] {
		node.Val = v
		if i == len(temp[1:])-1 {
			node.Next = nil
			break
		}
		node = node.Next
	}
}
