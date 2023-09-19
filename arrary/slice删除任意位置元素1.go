package main

import "fmt"

func main() {
	s1 := []int{1, 2, 4, 5, 6, 4}
	//showArray(s1)
	DeletedSlice(s1, 2)

	showSlice(s1)
}

// 向前移一位
func DeletedSlice(slice []int, index int) {
	for i := index; i < len(slice)-1; i++ {
		fmt.Println(slice)
		slice[i] = slice[i+1]
	}
}

func showSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
	}
}
