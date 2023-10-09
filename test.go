package main

import "fmt"

func main() {
	str := "Hello, World!"
	var strSlice []string // 声明一个字符串切片

	for _, char := range str { // 遍历字符串的每个字符
		strSlice = append(strSlice, string(char)) // 将每个字符转换为字符串，并添加到切片中
	}

	fmt.Println(strSlice) // 输出切片内容
}
