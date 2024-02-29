package main

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	for i := 0; i < n; i++ {
		nums1 = append(nums1, nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
}
