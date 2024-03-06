package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 3, 1, 1, 4}
	//fmt.Println(canJump(arr))
	fmt.Println(len(arr))
	fmt.Println(canJump(arr))
}

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		fmt.Println("start", i, v)
		if i > maxJump {
			return false
		}
		maxJump = canJumpTest(maxJump, i+v)
		fmt.Println("end", i, maxJump)
	}
	return true
}

func canJumpTest(a, b int) int {
	if a > b {
		return a
	}
	return b
}
