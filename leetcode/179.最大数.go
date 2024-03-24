package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{10, 2}
	fmt.Println(largestNumber(arr))
}

func largestNumber(nums []int) string {

	var res string
	for _, v := range nums {
		res += strconv.Itoa(v)
	}
	return res
}
