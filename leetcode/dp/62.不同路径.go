package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				result[i][j] = 1
				continue
			}
			result[i][j] = result[i][j-1] + result[i-1][j]
		}
	}
	return result[m-1][n-1]
}
