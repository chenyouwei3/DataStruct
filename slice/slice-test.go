package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} /* 创建切片 */
	/* 打印原始切片 */ /* [0 1 2 3 4 5 6 7 8]*/
	fmt.Println("numbers ==", numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/ /*[1 2 3]*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])
	/* 默认下限为 0*/ /* [0 1 2]*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/* 默认上限为 len(s)*/ /*[4 5 6 7 8]*/
	fmt.Println("numbers[4:] ==", numbers[4:])

}
