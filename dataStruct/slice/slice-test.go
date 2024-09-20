package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} /* 创建切片 */
	/* 打印原始切片 */ /* [0 1 2 3 4 5 6 7 8]*/
	fmt.Println("numbers ==", numbers)
	fmt.Println(numbers[0:1])
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/ /*[1 2 3]  */
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/ /* [0 1 2]*/ //不包含索引3
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/ /*[3 4 5 6 7 8]*/
	fmt.Println("numbers[3:] ==", numbers[3:])
}

func deletedSlice() {
	slice := []int{1, 2, 3, 4, 5}
	// 要删除的元素的索引
	index := 4
	// 使用切片的切片方式删除元素
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
}
