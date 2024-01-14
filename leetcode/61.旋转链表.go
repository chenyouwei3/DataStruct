package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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
	res := rotateRight(head, 2)
	for res != nil {
		fmt.Printf("->%d", res.Val)
		res = res.Next
	}
}

func rotateRight(head *ListNode, k int) *ListNode {
	pre := new(ListNode)
	cur, res := head, pre
	for i := 0; i < k; i++ {
		newnode := &ListNode{Val: cur.Val, Next: nil}
		rotateRightFZ(pre, newnode)
		cur, pre = cur.Next, pre.Next
	}
	res = res.Next
	fmt.Println("cur:", cur, "pre:", pre)
	for res != nil {
		fmt.Printf("->%d", res.Val)
		res = res.Next
	}
	fmt.Println("")
	for cur != nil {
		fmt.Printf("->%d", cur.Val)
		cur = cur.Next
	}
	//fmt.Println(head)
	//fmt.Println(head.Next)
	//fmt.Println(head.Next.Next)
	return cur
}
func rotateRightFZ(node0 *ListNode, p *ListNode) {
	Temp := node0.Next //存储中间变量
	p.Next = Temp      //把目标节点的后面连接原来的后面
	node0.Next = p     //把节点node0连接后面
}
