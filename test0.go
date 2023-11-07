package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID    int
	Image string
}

// 模拟处理任务
func process(t *Task) {
	// 模拟图片压缩处理，这里只是演示，并没有实际的图像处理逻辑
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d processed image: %s\n", t.ID, t.Image)
}

// 模拟发送任务到通道中
func sendWork(pending chan<- *Task) {
	tasks := []*Task{
		{ID: 1, Image: "image1.jpg"},
		{ID: 2, Image: "image2.jpg"},
		{ID: 3, Image: "image3.jpg"},
		{ID: 4, Image: "image4.jpg"},
	}
	for _, task := range tasks {
		pending <- task
	}
	//关闭通道,表示任务发送完成
	close(pending)
}

func main() {
	pending, done := make(chan *Task), make(chan *Task)
	go sendWork(pending) // 将有工作的任务放到频道上

	//开启任务
	for i := 0; i < 4; i++ { // 开始N次外出工作
		go Worker(pending, done)
	}
	consumeWork(done) // 继续处理已处理的任务
}

/* ---------------------------------------------------*/
// 工作协程处理任务
func Worker(pending <-chan *Task, done chan<- *Task) {
	for task := range pending {
		process(task)
		done <- task
	}
}

// 消费已处理的任务
func consumeWork(done <-chan *Task) {
	for task := range done {
		// 根据实际需求处理已完成的任务
		fmt.Printf("Consuming processed task %d\n", task.ID)
	}
}
