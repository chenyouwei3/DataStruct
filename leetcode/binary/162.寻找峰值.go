package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(arr))
}

func findPeakElement(nums []int) (idx int) {
	fmt.Println(idx, "gggg")
	for i, v := range nums {
		fmt.Println(i, v)
		if v > nums[idx] {
			fmt.Println("23131", i, v)
			idx = i
		}
	}
	return idx
}
