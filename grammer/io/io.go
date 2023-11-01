package main

import (
	"fmt"
	"os"
)

func main() {
	//创建
	f, err := os.Create("./grammer/32.txt")
	fmt.Println(f) // 打印文件指针
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := f.Write([]byte("hello"))
	fmt.Println("write number = ", n)
	if err != nil {
		fmt.Println("write err: ", err)
		return
	}
	f.Close()
}
