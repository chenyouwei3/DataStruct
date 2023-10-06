package main

import "fmt"

type Node struct {
	data interface{} // 节点存储的数据
	pre  *Node       // 指向前一个节点的指针
	next *Node       // 指向后一个节点的指针
}

// DoubleList 双向链表
type DoubleList struct {
	Head *Node
	Tail *Node
}

// PushTail 在链表末尾插入节点
func (list *DoubleList) PushTail(data interface{}) {
	newNode := &Node{
		data: data,
		pre:  nil,
		next: nil,
	}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.pre = list.Tail
		list.Tail.next = newNode
		list.Tail = newNode //更新节点
	}
}

// PushHead 在链表头部插入节点
func (list *DoubleList) PushHead(data interface{}) {
	newNode := &Node{
		data: data,
		pre:  nil,
		next: nil,
	}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.next = list.Head
		list.Head.pre = newNode
		list.Head = newNode //更新节点
	}
}

// Remove Remove删除指定节点
func (list *DoubleList) Remove(node *Node) {
	if node == list.Head {
		// 如果要删除的节点是头节点，则更新头节点为下一个节点，并设置前指针为nil
		list.Head = list.Head.next
		if list.Head != nil {
			list.Head.pre = nil
		}
	} else if node == list.Tail {
		// 如果要删除的节点是尾节点，则更新尾节点为前一个节点，并设置后指针为nil
		list.Tail = list.Tail.pre
		if list.Tail != nil {
			list.Tail.next = nil
		}
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
}

// Print 遍历链表并打印节点的值
func (list *DoubleList) Print() {
	node := list.Head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

func main() {
	dlist := DoubleList{}
	dlist.PushHead(3)
	dlist.PushHead(2)
	dlist.PushHead(1)
	dlist.PushTail(4)
	dlist.Print()
}
