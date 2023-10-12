package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// 初始化链表
	head := &ListNode{Val: 2}
	node1 := &ListNode{Val: 5}
	node2 := &ListNode{Val: 3}
	head.Next = node1
	node1.Next = node2

	// 初始化链表
	Head := &ListNode{Val: 5}
	Node1 := &ListNode{Val: 6}
	Node2 := &ListNode{Val: 4}
	Head.Next = Node1
	Node1.Next = Node2

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	charMap := make(map[int]int)
	for cur1 != nil {

	}
}
