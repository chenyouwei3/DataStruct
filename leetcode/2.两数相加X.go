package main

import "fmt"

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	head := &ListNode{Val: 2}
	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 3}
	head.Next = node1
	node1.Next = node2

	Head := &ListNode{Val: 5}
	Node1 := &ListNode{Val: 6}
	Node2 := &ListNode{Val: 4}
	//Node3 := &ListNode{Val: 1}
	Head.Next = Node1
	Node1.Next = Node2
	//Node2.Next = Node3
	//head := &ListNode{Val: 1}
	//node1 := &ListNode{Val: 2}
	//node2 := &ListNode{Val: 3}
	//node3 := &ListNode{Val: 4}
	//node4 := &ListNode{Val: 5}
	//node5 := &ListNode{Val: 6}
	//node6 := &ListNode{Val: 7}
	//node7 := &ListNode{Val: 8}
	////node8 := &ListNode{Val: 4}
	//head.Next = node1
	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = node4
	//node4.Next = node5
	//node5.Next = node6
	//node6.Next = node7
	////node7.Next = node8
	//Head := &ListNode{Val: 9}
	//Node1 := &ListNode{Val: 9}
	//Node2 := &ListNode{Val: 9}
	//Node3 := &ListNode{Val: 9}
	//Head.Next = Node1
	//Node1.Next = Node2
	//Node2.Next = Node3
	test := addTwoNumbers(head, Head)
	for test != nil {
		fmt.Printf("->%d", test.Val)
		test = test.Next
	}

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义一个虚拟头节点
	head := &ListNode{Val: 0}
	// 定义变量，其中 n1 和 n2 代表两个链表中的值，carry 用于存储进位信息，current 用于记录新链表当前节点
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10} // 计算当前位置上两数之和以及进位数
		current = current.Next                                // 将链表指针后移至下一个节点
		carry = (n1 + n2 + carry) / 10                        // 计算该位置上的进位数
		fmt.Println(carry)
	}
	return head.Next
}

func addTwoNumbersTest(l1 *ListNode, l2 *ListNode) *ListNode {
	boll := false
	if l1.Val == 0 && l2.Val == 0 && l1.Next == nil && l2.Next == nil {
		l3 := &ListNode{Val: 0}
		return l3
	}
	cur1, cur2 := l1, l2
	l3 := new(ListNode)
	cur := l3
	for cur1 != nil && cur2 != nil {
		total := cur1.Val + cur2.Val
		if boll == true {
			total = total + 1
		}
		if total >= 10 {
			boll = true
			total = total % 10
		} else {
			boll = false
		}
		newnode := &ListNode{Val: total}
		cur1, cur2 = cur1.Next, cur2.Next
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newnode
		if cur2 == nil {
			p := &ListNode{Val: 0}
			cur2 = p
		}
		if cur1 == nil {
			p := &ListNode{Val: 0}
			cur1 = p
		}
		if cur1.Val == 0 && cur2.Val == 0 && boll == false {
			if cur1.Next == nil && cur2.Next == nil {
				break
			}
		}
	}
	l4 := l3.Next
	return l4
}
