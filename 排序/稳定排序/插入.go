package main

import "fmt"

// 时间复杂度
// best      O(n)
// average   O(n^2)
// worst     O(n^2)

// 空间复杂度O(1)
func main() {
	arr := []int{9, 5, 1, 4, 8, 2, 7, 3, 6}
	insertsort(arr)
}

// 从第二个元素开始（假设第一个元素是已排序的）。
// 遍历每个元素，将其与前面已经排好序的元素进行比较。
// 将当前元素插入到前面已排序部分中的正确位置，以保证插入后的部分仍然是有序的。
// 重复上述步骤，直到最后一个元素被插入到正确位置。
func insertsort(arr []int) {
	// 从第二个元素开始，默认第一个元素已经排序
	for i := 1; i < len(arr); i++ {
		// 保存当前要插入的元素
		key := arr[i]
		// j 是已排序部分的最后一个元素的索引
		j := i - 1
		// 内层循环：将已排序部分中比 key 大的元素向后移动
		for ; j >= 0 && arr[j] > key; j-- {
			// 将元素向后移动一位
			arr[j+1] = arr[j]
		}
		// 将 key 插入到正确的位置
		arr[j+1] = key
	}

	// 打印排序后的数组
	fmt.Println(arr)
}
