package main

import "fmt"

// 时间复杂度
// best      O(n)
// average   O(n^2)
// worst     O(n^2)

// 空间复杂度O(1)
func main() {
	arr := []int{1, 8, 2, 3, 4, 6, 5, 9, 8, 7}
	fmt.Println("正常结果", bubbleSort(arr))
	fmt.Println("测试结果", bubbleSortTest(arr))
}

// 什么时候最快?
// 当输入的数据已经是正序时（都已经是正序了，我还要你冒泡排序有何用啊）。
//
// 什么时候最慢?
// 当输入的数据是反序时（写一个 for 循环反序输出数据不就行了，干嘛要用你冒泡排序呢，我是闲的吗）
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

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
