package main

import "fmt"

// ListNode 节点结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedListStacks 栈结构体
type LinkedListStacks struct {
	top *ListNode
}

// 入栈
func (s *LinkedListStacks) push(Val int) {
	newNode := &ListNode{
		Val:  Val,
		Next: s.top,
	}
	s.top = newNode
}

// Pop 出栈
func (s *LinkedListStacks) Pop() int {
	if s.IsEmpty() {
		fmt.Println("栈为空")
	}
	val := s.top.Val
	s.top = s.top.Next
	return val
}

// IsEmpty 判断栈是否为空
func (s *LinkedListStacks) IsEmpty() bool {
	return s.top == nil
}
func main() {
	stack := LinkedListStacks{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())
}
