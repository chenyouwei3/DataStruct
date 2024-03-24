package main

import "fmt"

func main() {
	arr := []int{2, 2, 1}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	buckets := make(map[int]int)
	for _, v := range nums {
		buckets[v]++
	}
	for i, v := range buckets {
		if v == 1 {
			return i
		}
	}
	return 0
}
