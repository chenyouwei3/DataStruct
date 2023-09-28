package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedListQueue 队列结构体
type LinkedListQueue struct {
	front *ListNode
	rear  *ListNode
}

// Enqueue 入列操作
func (q *LinkedListQueue) Enqueue(val int) {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		//连接到尾部
		q.rear.Next = newNode
		//更新节点
		q.rear = newNode
	}
}

// Dequeue 出队列操作
func (q *LinkedListQueue) Dequeue() int {
	if q.front == nil {
		fmt.Println("该队列为空")
	}
	val := q.front.Val
	q.front = q.front.Next
	if q.front == nil {
		q.rear = nil
	}
	return val
}

// IsEmpty 判断队列是否为空
func (q *LinkedListQueue) IsEmpty() bool {
	return q.front == nil
}

func main() {
	queue := LinkedListQueue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.IsEmpty())
}
