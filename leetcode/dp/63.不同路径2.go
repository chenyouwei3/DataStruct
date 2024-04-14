package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid[0]), len(obstacleGrid)
	result := make([][]int, n)
	for i := range obstacleGrid {
		result[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
				result[i][j] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				continue
			}
			if i == 0 {
				result[i][j] += result[i][j-1]
				continue
			}
			if j == 0 {
				result[i][j] += result[i-1][j]
				continue
			}
			result[i][j] = result[i][j-1] + result[i-1][j]
		}
	}
	return result[n-1][m-1]
}
