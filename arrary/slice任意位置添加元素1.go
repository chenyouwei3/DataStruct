package main

import "fmt"

func main() {

	var slice []int
	slice = []int{1, 3, 5, 4, 5}
	//showArray(arr)
	slice = append(slice, 0)
	InsertAtIndexSlice(slice, 8, 3)
	fmt.Println("--------")
	showSlice(slice)
}

// 向后移一位
func InsertAtIndexSlice(slice []int, data int, index int) {
	for i := len(slice) - 1; i > index; i-- {
		slice[i] = slice[i-1]
	}
	slice[index] = data
}

//func showSlice(slice []int) {
//	for i := 0; i < len(slice); i++ {
//		fmt.Print(slice[i])
//	}
//}
