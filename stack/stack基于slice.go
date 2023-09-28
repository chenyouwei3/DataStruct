package main

import "fmt"

//stack=一种受限制的数组或者链表

func main() {
	//初始化stack
	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	stack = append(stack, 4)
	stack = append(stack, 5)
	stack = append(stack, 6)
	//访问栈顶元素
	peek := stack[len(stack)-1]

	//元素出栈
	pop := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Print(pop)
	fmt.Print(stack)
	fmt.Println(peek)
}
