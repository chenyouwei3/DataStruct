package main

func main() {
	high := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea(high)
}

// 双指针
func maxArea(height []int) int {
	left, area, res := 0, 0, 0 // 初始化左边界、面积和结果为0
	right := len(height) - 1   // 初始化右边界为数组最后一个元素的索引

	for left < right { // 循环直到左边界达到右边界
		area = (right - left) * minTest(height[left], height[right]) // 计算当前容器的面积
		if height[left] < height[right] {                            // 如果左边界的高度小于右边界的高度
			left++ // 左边界向右移动
		} else {
			right-- // 否则，右边界向左移动
		}
		res = maxTest(res, area) // 更新最大面积
	}
	return res // 返回最大面积
}

// 暴力法
func maxAreaTest(height []int) int {
	area := 0
	Area1 := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area = (j - i) * minTest(height[i], height[j])
			Area1 = maxTest(area, Area1)
		}
	}
	return Area1
}
func maxTest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minTest(a, b int) int {
	if a < b {
		return a
	}
	return b
}
