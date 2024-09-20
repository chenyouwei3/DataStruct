package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	left, right, sum, maxCount := 0, 0, a[0], 0
	for right < n {
		if sum <= m {
			maxCount = max(maxCount, right-left+1)
			right++
			if right < n {
				sum += a[right]
			}
		} else {
			sum -= a[left]
			left++
		}
	}
	fmt.Println(maxCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
