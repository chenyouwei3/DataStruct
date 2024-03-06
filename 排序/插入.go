package main

import "fmt"

func main() {
	arr := []int{9, 5, 1, 4, 8, 2, 7, 3, 6}
	insertsort(arr)
	testInsertSort(arr)
}

func insertsort(arr []int) {
	// 从第二个元素开始遍历，因为第一个元素已经被认为是有序的
	for i := 1; i < len(arr); i++ {
		// 将当前元素作为关键元素
		key := arr[i]
		// 将 j 初始化为当前元素的前一个索引
		j := i - 1
		// 向前遍历已排序部分，直到找到一个元素小于或等于关键元素的位置
		// 在遍历的过程中，不断将比关键元素大的元素向右移动一个位置
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j] // 将元素向右移动一个位置
		}
		// 插入关键元素到正确的位置（j+1）
		arr[j+1] = key
	}
	// 打印排序后的数组
	fmt.Println(arr)
}

func test(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}

func testInsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}
