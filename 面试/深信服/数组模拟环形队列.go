package main

import (
	"errors"
	"fmt"
)

// CircularQueue 定义了环形队列的结构
type CircularQueue struct {
	data  []int
	front int
	rear  int
	size  int
	cap   int
}

// NewCircularQueue 创建一个新的环形队列，capacity 为队列的容量
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, capacity),
		front: 0,
		rear:  0,
		size:  0,
		cap:   capacity,
	}
}

// Enqueue 将元素添加到队列的尾部
func (q *CircularQueue) Enqueue(value int) error {
	if q.IsFull() {
		return errors.New("circular queue is full")
	}
	q.data[q.rear] = value
	q.rear = (q.rear + 1) % q.cap
	q.size++
	return nil
}

// Dequeue 从队列的头部删除并返回一个元素
func (q *CircularQueue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("circular queue is empty")
	}
	value := q.data[q.front]
	q.front = (q.front + 1) % q.cap
	q.size--
	return value, nil
}

// FrontElement 返回队列头部的元素，不删除
func (q *CircularQueue) FrontElement() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("circular queue is empty")
	}
	return q.data[q.front], nil
}

// RearElement 返回队列尾部的元素，不删除
func (q *CircularQueue) RearElement() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("circular queue is empty")
	}
	return q.data[(q.rear-1+q.cap)%q.cap], nil
}

// IsEmpty 检查队列是否为空
func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull 检查队列是否已满
func (q *CircularQueue) IsFull() bool {
	return q.size == q.cap
}

// Size 返回队列的大小
func (q *CircularQueue) Size() int {
	return q.size
}

// PrintQueue 打印队列中的所有元素
func (q *CircularQueue) PrintQueue() {
	if q.IsEmpty() {
		fmt.Println("Circular Queue is empty")
		return
	}
	fmt.Print("Circular Queue: ")
	for i := 0; i < q.size; i++ {
		index := (q.front + i) % q.cap
		fmt.Print(q.data[index], " ")
	}
	fmt.Println()
}

func main() {
	cq := NewCircularQueue(5) // 创建一个容量为 5 的环形队列

	// 入队操作
	if err := cq.Enqueue(10); err == nil {
		fmt.Println("Enqueued:", 10)
	}
	if err := cq.Enqueue(20); err == nil {
		fmt.Println("Enqueued:", 20)
	}
	if err := cq.Enqueue(30); err == nil {
		fmt.Println("Enqueued:", 30)
	}
	cq.PrintQueue() // 输出: Circular Queue: 10 20 30

	// 出队操作
	value, err := cq.Dequeue()
	if err == nil {
		fmt.Println("Dequeued:", value) // 输出: Dequeued: 10
	}
	cq.PrintQueue() // 输出: Circular Queue: 20 30

	// 再次入队
	if err := cq.Enqueue(40); err == nil {
		fmt.Println("Enqueued:", 40)
	}
	if err := cq.Enqueue(50); err == nil {
		fmt.Println("Enqueued:", 50)
	}
	if err := cq.Enqueue(60); err == nil {
		fmt.Println("Enqueued:", 60)
	}
	cq.PrintQueue() // 输出: Circular Queue: 20 30 40 50 60

	// 尝试入队到满的队列
	if err := cq.Enqueue(70); err != nil {
		fmt.Println("Error:", err) // 输出: Error: circular queue is full
	}

	// 查看队列前端和后端元素
	front, err := cq.FrontElement()
	if err == nil {
		fmt.Println("Front Element:", front) // 输出: Front Element: 20
	}

	rear, err := cq.RearElement()
	if err == nil {
		fmt.Println("Rear Element:", rear) // 输出: Rear Element: 60
	}

	// 队列大小
	fmt.Println("Size of circular queue:", cq.Size()) // 输出: Size of circular queue: 5

	// 出队所有元素
	for !cq.IsEmpty() {
		value, _ = cq.Dequeue()
		fmt.Println("Dequeued:", value)
	}
	cq.PrintQueue() // 输出: Circular Queue is empty
}
