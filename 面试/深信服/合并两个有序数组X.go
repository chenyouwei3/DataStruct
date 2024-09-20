package main

import "fmt"

func main() {
	var test1 = []int{1, 2, 3, 0, 0, 0}
	var test2 = []int{2, 5, 6}
	merge(test1, 3, test2, 3)
	fmt.Println(test1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1 //最后一个位置
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
	fmt.Println(nums1) // 输出合并后的 nums1
	fmt.Println(nums2) // 输出 nums2 的内容（没有变化）
}
