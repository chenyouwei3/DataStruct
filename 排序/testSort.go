package main

import (
	"fmt"
)

func main() {
	arr := []int{0, 3, 7, 1, 6, 4, 9, 8, 5, 2, 4, 7, 2}
	fmt.Println("二分查找:", BinarySeach(arr, 7))
	testSort0(arr)
	testSort1(arr)
	testSort2(arr)
	fmt.Println(testSort3(arr))
}

// 二分
func BinarySeach(nums []int, taregt int) int {
	min, max := 0, len(nums)-1
	for min < max {
		mid := (max + min) / 2
		if nums[mid] == taregt {
			return mid
		} else if nums[mid] > taregt {
			max--
		} else if nums[mid] < taregt {
			min--
		}
	}
	return -1
}

// 冒泡
func testSort0(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 选择
func testSort1(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}

// 插入
func testSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}

// 快排
func testSort3(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	proindex := len(arr) / 2
	pro := arr[proindex]
	var left, right []int
	for i, v := range arr {
		if i == proindex {
			continue
		}
		if v <= pro {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = testSort3(left)
	right = testSort3(right)
	return append(append(left, pro), right...)
}
