package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1 // 或者根据需求返回其他错误码或错误值
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue()) // 输出：1
	fmt.Println(queue.Dequeue()) // 输出：2
	fmt.Println(queue.Dequeue()) // 输出：3
	fmt.Println(queue.Dequeue()) // 输出：-1，队列为空
}
