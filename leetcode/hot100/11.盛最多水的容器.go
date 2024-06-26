package main

import "fmt"

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(maxArea(arr))
}
func maxArea(height []int) int {
	left, area, res := 0, 0, 0 // 初始化左边界、面积和结果为0
	right := len(height) - 1   // 初始化右边界为数组最后一个元素的索引
	for left < right {         // 循环直到左边界达到右边界
		area = (right - left) * minAreaTest(height[left], height[right]) // 计算当前容器的面积
		if height[left] < height[right] {                                // 如果左边界的高度小于右边界的高度
			left++ // 左边界向右移动
		} else {
			right-- // 否则，右边界向左移动
		}
		res = maxAreaTest(res, area) // 更新最大面积
	}
	return res // 返回最大面积
}

func maxAreaTest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minAreaTest(a, b int) int {
	if a < b {
		return a
	}
	return b
}
