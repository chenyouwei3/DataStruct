package main

import "fmt"

func main() {
	arr := []int{1000000000, 1000000000, -1000000000, -1000000000, -1000000000}
	fmt.Println(majorityElement(arr))
}

func majorityElement(nums []int) int {
	mmap := make(map[int]int)
	res := 0
	for i := 0; i < len(nums); i++ {
		mmap[nums[i]]++
		if mmap[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return res
}
