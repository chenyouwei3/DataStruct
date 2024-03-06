package main

import "fmt"

func main() {

	fmt.Println(1 % 2)

}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}
	maxLength := 0
	for i, v := range nums {
		if i > maxLength {
			return false
		}
		maxLength = max(maxLength, i+v)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
