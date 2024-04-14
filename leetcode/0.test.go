package main

import "fmt"

func main() {
	str := "abcde"
	str1 := "ace"
	fmt.Println(longestCommonSubsequence(str, str1))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	buckets := make([]int32, 0)
	var temp1, temp2 string
	var res int
	if len(text1) > len(text2) {
		temp1, temp2 = text1, text2
	} else {
		temp1, temp2 = text2, text1
	}
	for _, v := range temp1 {
		buckets = append(buckets, v)
	}
	for _, v := range temp2 {

	}
	return res
}
