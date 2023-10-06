package main

import "fmt"

type Deque struct {
	data []interface{}
}

// PushHead 队列头插
func (db *Deque) PushHead(value interface{}) {
	db.data = append([]interface{}{value}, db.data...)
}

func (db *Deque) PushTail(value interface{}) {
	db.data = append(db.data, value)
}

// PopTail 删除后面元素
func (db *Deque) PopTail() interface{} {
	if db.IsEmpty() {
		return nil
	}
	value := db.data[len(db.data)-1]
	db.data = db.data[:len(db.data)-1]
	return value
}

// PopHead 删除前面元素
func (db *Deque) PopHead() interface{} {
	if db.IsEmpty() {
		return nil
	}
	value := db.data[0]
	db.data = db.data[1:]
	return value
}

func main() {
	dq := Deque{}

	dq.PushHead(2)
	dq.PushHead(1)
	dq.PushTail(3)

	fmt.Println("Front:", dq.Front()) // 输出：Front: 1
	fmt.Println("Back:", dq.Back())   // 输出：Back: 2
	fmt.Println("Size:", dq.Size())   // 输出：Size: 3

	dq.PopHead()
	dq.PopTail()
	dq.PopTail()
	fmt.Println("IsEmpty:", dq.IsEmpty()) // 输出：IsEmpty: false
}

// Front 获取队列前端元素
func (db *Deque) Front() interface{} {
	if db.IsEmpty() {
		return nil
	}
	return db.data[0]
}

// Back 获取队列后端元素
func (db *Deque) Back() interface{} {
	if db.IsEmpty() {
		return nil
	}
	return db.data[len(db.data)-1]
}

// IsEmpty 判断队列是否为空
func (db *Deque) IsEmpty() bool {
	return len(db.data) == 0
}

// Size 获取队列大小
func (db *Deque) Size() int {
	return len(db.data)
}
