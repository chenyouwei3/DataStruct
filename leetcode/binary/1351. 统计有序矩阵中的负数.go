package main

import "fmt"

func main() {
	letters := []byte{'c', 'f', 'j'}
	fmt.Println(letters)
	fmt.Println(nextGreatestLetter(letters, 'a'))
}

func countNegatives(grid [][]int) int {
	var count int
	for _, r := range grid {
		f, e, ans := 0, len(r)-1, -1
		for f <= e {
			mid := f + (e-f)/2
			if r[mid] >= 0 {
				f = mid + 1
			} else {
				ans = mid
				e = mid - 1
			}
		}
		if ans != -1 {
			count = count + len(r) - ans
		}
	}
	return count
}
