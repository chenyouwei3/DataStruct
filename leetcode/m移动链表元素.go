package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	head := &ListNode{Val: 1}
	node1 := &ListNode{Val: 9}
	node2 := &ListNode{Val: 3}
	node3 := &ListNode{Val: 9}
	node4 := &ListNode{Val: 5}
	node5 := &ListNode{Val: 7}
	node6 := &ListNode{Val: 98}
	node7 := &ListNode{Val: 90}
	//node8 := &ListNode{Val: 4}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	//node7.Next = node8
	Head := &ListNode{Val: 1}
	Node1 := &ListNode{Val: 0}
	Node2 := &ListNode{Val: 9}
	Node3 := &ListNode{Val: 3}
	Head.Next = Node1
	Node1.Next = Node2
	Node2.Next = Node3
	//removeElements(Head, 9)

}

//func removeElements(head *ListNode, val int) *ListNode {
//
//}
