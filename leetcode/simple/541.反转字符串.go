package main

import "fmt"

func main() {
	s := "chenyouwei"
	fmt.Println("res", reverseStr(s, 3))
}

func reverseStr(s string, k int) string {
	length := len(s) - 1
	fmt.Println(length)
	result := make([]byte, length+1)
	for i := 0; i <= k; i++ {
		fmt.Println("adc", string(s[length-i]), i)
		result[i] = s[length-i]
	}
	fmt.Println(result)
	fmt.Println(string(result))
	return ""
}
