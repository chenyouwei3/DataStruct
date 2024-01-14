package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	close(ch) // 关闭通道
	mi := <-ch
	fmt.Println(mi)

}
