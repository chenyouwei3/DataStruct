package main

func main() {
	high := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(high)
}

func maxArea(height []int) {
	left, maxS := 0, 0
	right := len(height) - 1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
		maxS = max1(maxS, area)
	}
}

// 时间复杂度n 2
func maxAreaTest(height []int) int {
	area := 0
	Area1 := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area = (j - i) * min(height[i], height[j])
			Area1 = max1(area, Area1)
		}
	}
	return Area1
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
