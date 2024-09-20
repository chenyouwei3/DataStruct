package main

import (
	"fmt"
	"sync"
)

func printNumber(n int) {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(n int) {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-c1
			fmt.Println("xixi", i)
			c2 <- true
		}
	}(n)

	go func(n int) {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-c2
			fmt.Println(i)
			c1 <- true
		}
	}(n)

	// 启动第一个goroutine，因为从1开始，所以先启动打印奇数的协程
	c2 <- true

	wg.Wait()
}

func main() {
	printNumber(10)
}
