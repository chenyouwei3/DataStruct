package main

import (
	"fmt"
	"time"
)

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("收到消息:", num)
	}
}

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Producer: sending", i)
		ch <- i // 发送数据到通道
		time.Sleep(time.Second)
	}
	close(ch) // 关闭通道
}

func main() {
	ch := make(chan int) // 创建一个整数类型的通道
	// 启动生产者goroutine
	go producer(ch)
	// 启动消费者goroutine
	go consumer(ch)
	// 等待一段时间，确保生产者和消费者有足够的时间执行
	time.Sleep(7 * time.Second)
}
