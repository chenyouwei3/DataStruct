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
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]
	// 定义左右两个子数组
	var left, right []int
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
	left = quickSort(left)
	right = quickSort(right)
	sorted := append(append(left, pivot), right...)
	return sorted
}
