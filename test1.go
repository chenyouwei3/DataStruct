package main

import "fmt"

func main() {
	// 定义一个切片
	slice := []int{1, 2, 3, 4, 5}

	// 要删除的元素的索引
	index := 2

	// 使用切片的切片方式删除元素
	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println(slice) // [1 2 4 5]
}