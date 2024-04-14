package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(arr))
}

func trap(height []int) int {
	area := 0
	for i := 1; i < len(height)-2; i++ {
		if height[i-1] <= height[i] && height[i] <= height[i+1] {
			area += trapMin(height[i-1], height[i+1])
		}
		fmt.Println(area)
	}
	return 0
}

func trapMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}
