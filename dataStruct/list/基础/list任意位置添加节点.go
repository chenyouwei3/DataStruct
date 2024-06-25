package main

import "fmt"

func main() {
	var head = new(ListNode)
	head.data = 1
	var node1 = new(ListNode)
	node1.data = 2
	var node2 = new(ListNode)
	node2.data = 3

	head.next = node1
	node1.next = node2
	newnode := &ListNode{data: 8}
	InsertAtIndexList1(node1, newnode)
	//InsertAtIndexListTest(head, 8, 1)
	for head != nil {
		fmt.Printf("%d->", head.data)
		head = head.next
	}

}

// InsertAtIndexList1 插在0的后面
func InsertAtIndexList1(node0 *ListNode, p *ListNode) {
	Temp := node0.next //存储中间变量
	p.next = Temp      //把目标节点的后面连接原来的后面
	node0.next = p     //把节点node0连接后面
}

func InsertAtIndexListTest(head *ListNode, data int, index int) {
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
	//把目标节点位置后面的节点往后面移动
	for cur != nil && cout < index {
		//往后面遍历
		Temp := cur.next //存储中间变量
		pre = cur        //移动节点
		cur = Temp       //移动节点
		//
		cout++
		fmt.Println(pre, "pre") //1
		fmt.Println(cur, "cur") //2
	}
	//添加节点
	pre.next = newNode
	newNode.next = cur
	//遍历节点
	for head != nil {
		fmt.Printf("-->%d", head.data)
		head = head.next
	}

}
