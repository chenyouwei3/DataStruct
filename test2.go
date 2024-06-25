package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 初始创建一个容量为2的切片
	slice := make([]int, 2, 2)
	fmt.Printf("Initial slice: len=%d cap=%d\n", len(slice), cap(slice))
	fmt.Println("地址", unsafe.Pointer(&slice))
	// 添加元素以触发扩容
	for i := 1; i <= 5; i++ {
		slice = append(slice, i)
		fmt.Printf("After appending %d: len=%d cap=%d\n", i, len(slice), cap(slice))
		fmt.Println("地址", unsafe.Pointer(&slice))
	}
}
