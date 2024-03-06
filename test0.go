package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 随机选择基准值
	pivotIndex := 4
	pivot := arr[pivotIndex]
	// 将基准值移到数组末尾
	arr[pivotIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotIndex]
	// 定义左右两个子数组
	var left, right []int
	// 遍历数组，将比基准值小的元素放入left数组，将比基准值大的元素放入right数组
	for _, v := range arr[:len(arr)-1] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	// 递归地对左右子数组进行快速排序
	left = quickSort(left)
	right = quickSort(right)
	// 合并左右子数组和基准值
	sorted := append(append(left, pivot), right...)
	return sorted
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Original array:", arr)
	sorted := quickSort(arr)
	fmt.Println("Sorted array:", sorted)
}
