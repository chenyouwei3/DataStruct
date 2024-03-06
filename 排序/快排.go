package main

import "fmt"

func main() {
	arr := []int{9, 5, 1, 4, 8, 2, 7, 3, 6}
	fmt.Println(quickSort(arr))
}
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	// 选择第一个元素作为基准值
	pivotIndex := len(arr)
	pivot := arr[pivotIndex]
	// 定义左右两个子数组
	var left, right []int
	// 遍历数组，将比基准值小的元素放入left数组，将比基准值大的元素放入right数组
	for i, v := range arr {
		if i == pivotIndex {
			continue // 跳过基准值的比较
		}
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
