package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(arr))
}

func findDuplicate(nums []int) int {
	bucket := make(map[int]int)
	for _, num := range nums {
		bucket[num]++
		if bucket[num] == 2 {
			return num
		}
	}
	return 0
}
