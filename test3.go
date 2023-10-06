package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s[0:])
	fmt.Println(s[1:])
	fmt.Println(s[2:])
	fmt.Println(s[3:])
	fmt.Println(s[4:])
	fmt.Println(s[5:])
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i:])
	}
}
