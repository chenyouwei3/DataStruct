package main

import "fmt"

func main() {
	arr := []int{0, 3, 7, 1, 6, 4, 9, 8, 5, 2, 4, 7, 2}
	//TestSort0(arr)
	Test2Sort(arr)
}

func TestSort0(arr []int) {
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

func TestSort1(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] { //要是值比j大
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	fmt.Println(arr)
}

func Test2Sort(arr []int) {
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
