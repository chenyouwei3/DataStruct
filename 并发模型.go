package main

import (
	"bufio"
	"fmt"
	"log"
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
		go worker(taskQueue, done, &wg)
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
func worker(taskQueue <-chan Task, done chan<- bool, wg *sync.WaitGroup) {
	for task := range taskQueue {
		processTask(task.Address, task.ProcessFunc)
	}
	wg.Done()
}

// 模拟处理任务
func processTask(Address string, Func func(conn net.Conn)) {
	listen, err := net.Listen("tcp", Address) //代表监听的地址端口
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

func Test1(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(conn)
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte               // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:])   // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {                 // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		recvStr := string(buf[:n])              // 将读取的字节转换为字符串
		fmt.Println("收到client端发来的数据：", recvStr) // 打印接收到的字符串数据
		conn.Write([]byte(recvStr))             // 将接收到的数据通过连接发送回客户端
	}
	fmt.Println(conn, "1")
}

func Test2(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte               // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:])   // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {                 // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		recvStr := string(buf[:n])              // 将读取的字节转换为字符串
		fmt.Println("收到client端发来的数据：", recvStr) // 打印接收到的字符串数据

		conn.Write([]byte(recvStr)) // 将接收到的数据通过连接发送回客户端
	}
	fmt.Println(conn, "2")
}

func Test3(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn) // 创建一个带缓冲的读取器，用于从连接中读取数据
		var buf [128]byte               // 创建一个长度为 128 的字节数组，用于存储读取的数据
		n, err := reader.Read(buf[:])   // 从读取器中读取数据，并将数据存储到 buf 中，同时返回读取的字节数和可能的错误
		if err != nil {                 // 若读取过程中发生错误
			fmt.Println("read from client failed, err:", err) // 打印错误信息
			break                                             // 结束循环，退出处理函数
		}
		recvStr := string(buf[:n])              // 将读取的字节转换为字符串
		fmt.Println("收到client端发来的数据：", recvStr) // 打印接收到的字符串数据

		conn.Write([]byte(recvStr)) // 将接收到的数据通过连接发送回客户端
	}
	fmt.Println(conn, "3")
}
