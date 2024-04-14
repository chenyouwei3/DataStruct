package main

func singleNonDuplicate(nums []int) int {
	buckets := make(map[int]int)
	for _, num := range nums {
		buckets[num]++
	}
	for key, value := range buckets {
		if value == 1 {
			return key
		}
	}
	return 0
}
