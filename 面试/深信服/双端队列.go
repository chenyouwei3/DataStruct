package main

import (
	"errors"
	"fmt"
)

// Deque 定义了双端队列的结构
type Deque struct {
	data []int
}

// NewDeque 创建一个新的双端队列
func NewDeque() *Deque {
	return &Deque{
		data: []int{},
	}
}

// PushFront 在队列的前端插入一个元素
func (d *Deque) PushFront(value int) {
	d.data = append([]int{value}, d.data...)
}

// PushBack 在队列的后端插入一个元素
func (d *Deque) PushBack(value int) {
	d.data = append(d.data, value)
}

// PopFront 从队列的前端删除并返回一个元素
func (d *Deque) PopFront() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	value := d.data[0]
	d.data = d.data[1:]
	return value, nil
}

// PopBack 从队列的后端删除并返回一个元素
func (d *Deque) PopBack() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	value := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]
	return value, nil
}

// Front 返回队列前端的元素，不删除
func (d *Deque) Front() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	return d.data[0], nil
}

// Back 返回队列后端的元素，不删除
func (d *Deque) Back() (int, error) {
	if d.IsEmpty() {
		return 0, errors.New("deque is empty")
	}
	return d.data[len(d.data)-1], nil
}

// IsEmpty 检查队列是否为空
func (d *Deque) IsEmpty() bool {
	return len(d.data) == 0
}

// Size 返回队列的大小
func (d *Deque) Size() int {
	return len(d.data)
}

// PrintDeque 打印队列中的所有元素
func (d *Deque) PrintDeque() {
	fmt.Println("Deque:", d.data)
}

func main() {
	deque := NewDeque()

	// 在前端插入元素
	deque.PushFront(10)
	deque.PushFront(20)
	deque.PrintDeque() // 输出: Deque: [20 10]

	// 在后端插入元素
	deque.PushBack(30)
	deque.PushBack(40)
	deque.PrintDeque() // 输出: Deque: [20 10 30 40]

	// 删除前端元素
	value, err := deque.PopFront()
	if err == nil {
		fmt.Println("Popped from front:", value) // 输出: Popped from front: 20
	}
	deque.PrintDeque() // 输出: Deque: [10 30 40]

	// 删除后端元素
	value, err = deque.PopBack()
	if err == nil {
		fmt.Println("Popped from back:", value) // 输出: Popped from back: 40
	}
	deque.PrintDeque() // 输出: Deque: [10 30]

	// 获取前端和后端元素
	front, _ := deque.Front()
	back, _ := deque.Back()
	fmt.Println("Front element:", front) // 输出: Front element: 10
	fmt.Println("Back element:", back)   // 输出: Back element: 30

	// 队列大小
	fmt.Println("Size of deque:", deque.Size()) // 输出: Size of deque: 2
}
