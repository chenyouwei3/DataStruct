package main

import "fmt"

func main() {
	arr := []int{6, 7, 5, 1, 6, 4, 9, 8, 5, 2, 4, 7, 2}
	//fmt.Println(selectSort(arr))
	selectsort(arr)
}
func selectsort(arr []int) {
	// 获取数组长度
	length := len(arr)
	// 外层循环，从第一个元素到倒数第二个元素
	for i := 0; i < length; i++ {
		// 初始化最小值的索引为当前外层循环的索引
		min := i
		// 内层循环，从当前元素的下一个位置开始到数组末尾
		for j := i + 1; j < length; j++ {
			// 如果找到比当前最小值还小的元素，则更新最小值的索引
			if arr[min] > arr[j] {
				min = j
			}
		}
		// 将当前外层循环索引位置的元素与最小值位置的元素进行交换
		arr[i], arr[min] = arr[min], arr[i]
	}
	// 输出排序后的数组
	fmt.Println(arr)
}
