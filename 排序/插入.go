package main

import "fmt"

func main() {
	arr := []int{9, 5, 1, 4, 8, 2, 7, 3, 6}
	insertsort(arr)
}

func insertsort(arr []int) {
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
