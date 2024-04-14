package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(sort.SearchInts(slice, 3))
	fmt.Println(sort.SearchInts(slice, 6))
	fmt.Println(test(slice))
}

func test(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		index := sort.SearchInts(dp, num)
		fmt.Println(dp)
		if index == len(dp) {
			dp = append(dp, num)
		} else {
			dp[index] = num
		}
	}
	return len(dp)
}
