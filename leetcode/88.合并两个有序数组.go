package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 2)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for _, n1 := range nums1 {
		temp := &nums2
		for j, n2 := range nums2 {
			if n1 > n2 {
				insert := []int{n1}
				*temp = append(temp[:j], append(insert, temp[j:]...)...)
			}
		}

	}
	fmt.Println(nums1)
}
