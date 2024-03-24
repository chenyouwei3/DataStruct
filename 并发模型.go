package main

import (
	"fmt"
	"net"
	"sync"
)

type Task struct {
	Address     string
	ProcessFunc func(conn net.Conn)
}

// 模拟发送任务到通道中
func sendTasks(taskQueue chan<- Task) {
	SliceTask := []Task{
		{
			Address:     "0.0.0.0:8004",
			ProcessFunc: Test1,
		},
		{
			Address:     "0.0.0.0:8010",
			ProcessFunc: Test2,
		},
		{
			Address:     "0.0.0.0:8019",
			ProcessFunc: Test3,
		},
	}
	for _, task := range SliceTask {
		taskQueue <- task
	}
	close(taskQueue)
}
func main() {
	taskQueue := make(chan Task) // 创建任务队列
	done := make(chan bool)      // 创建完成信号通道
	workerCount := 3             // 设置工作 goroutine 的数量
	go sendTasks(taskQueue)      // 启动生产任务的 goroutine
	// 启动一定数量的工作 goroutine 来处理任务
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i := 0; i < workerCount; i++ {
		go worker(taskQueue, &wg)
	}
	// 等待工作 goroutine 完成处理
	go func() {
		wg.Wait()
		done <- true
	}()

	<-done // 等待完成信号
	fmt.Println("所有任务处理完成")
}

// 工作协程处理任务
func worker(taskQueue <-chan Task, wg *sync.WaitGroup) {
	for task := range taskQueue {
		processTask(task, task.ProcessFunc)
	}
	wg.Done()
}

// 模拟处理任务
func processTask(task Task, Func func(conn net.Conn)) {
	listen, err := net.Listen("tcp", task.Address) //代表监听的地址端口
	defer listen.Close()
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	fmt.Println("正在等待建立连接.....", listen.Addr())
	for { //这个for循环的作用是可以多次建立连接
		conn, err := listen.Accept() //请求建立连接，客户端未连接就会在这里一直等待
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Println("连接建立成功.....")
		go Func(conn)
	}
}
