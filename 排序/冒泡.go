package main

import "fmt"

func main() {
	arr := []int{1, 8, 2, 3, 4, 6, 5, 9, 8, 7}
	fmt.Println(bubbleSortTest(arr))
}

func bubbleSort(arr []int) {
	// 获取数组长度
	length := len(arr)
	// 外层循环，从第一个元素到倒数第二个元素
	for i := 0; i < length; i++ {
		// 内层循环，从第一个元素到倒数第二个未排序的元素
		for j := 0; j < length-1-i; j++ {
			// 比较相邻的两个元素，如果顺序不正确则交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// 输出排序后的数组
	fmt.Println(arr)
}

func bubbleSortTest(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
