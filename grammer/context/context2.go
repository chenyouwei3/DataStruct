package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

//context.Background() 用于创建空的父级上下文，作为顶级上下文使用，不包含任何附加的值或取消操作。
//context.TODO() 用于表示一种“无事可做”的状态，用作临时占位符上下文
//一般不应该在实际业务逻辑中使用，而是应该尽快替换为具体的上下文构造函数。

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

}
