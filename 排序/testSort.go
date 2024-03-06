package main

import "fmt"

func main() {
	arr := []int{0, 3, 7, 1, 6, 4, 9, 8, 5, 2, 4, 7, 2}
	testSort0(arr)
	testSort1(arr)
	testSort2(arr)
	//fmt.Println(testSort3(arr))
}

// 冒泡
func testSort0(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
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
		arr[min], arr[i] = arr[i], arr[min]
	}
	fmt.Println(arr)
}

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

func testSort3(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}
	//基准值
	provIndex := len(arr) - 1
	index := arr[provIndex]
	var left, right []int
	for i, v := range arr {
		if i == index {
			continue
		}
		if v <= index {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = testSort3(left)
	right = testSort3(right)
	return nil
}
