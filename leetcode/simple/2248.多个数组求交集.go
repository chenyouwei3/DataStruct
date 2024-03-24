package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := make([][]int, 3)
	slice[0] = []int{3, 1, 2, 4, 5}
	slice[1] = []int{1, 2, 3, 4}
	slice[2] = []int{3, 4, 5, 6}
	fmt.Println(intersection(slice))
}

func intersection(nums [][]int) []int {
	length := len(nums)
	buckets := make(map[int]int)
	for _, v := range nums {
		for _, value := range v {
			buckets[value]++
		}
	}
	var res []int
	for key, v := range buckets {
		if v == length {
			res = append(res, key)
		}
	}
	sort.Ints(res)
	return res
}
