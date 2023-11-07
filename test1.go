package main

import (
	"fmt"
	"sync"
	"time"
)

type Task0 struct {
	ID    int
	Image string
}

func main() {
	pending, done := make(chan *Task0), make(chan *Task0)
	var wg sync.WaitGroup
	wg.Add(1)
	go sendWork0(pending)    // 将有工作的任务放到频道上
	for i := 0; i < 4; i++ { // 开始N次外出工作
		go Worker0(pending, done, &wg)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
	consumeWork0(done) // 继续处理已处理的任务

}

// 模拟发送任务到通道中
func sendWork0(pending chan<- *Task0) {
	tasks := []*Task0{
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

// 工作协程处理任务
func Worker0(pending <-chan *Task0, done chan<- *Task0, wg *sync.WaitGroup) {
	for task := range pending {
		process0(task)
		done <- task
	}
	wg.Done()
}

// 模拟处理任务
func process0(t *Task0) {
	// 模拟图片压缩处理，这里只是演示，并没有实际的图像处理逻辑
	time.Sleep(1 * time.Second)
	fmt.Printf("正在处理第 %d 个任务 image: %s\n", t.ID, t.Image)
}

// 消费已处理的任务
func consumeWork0(done <-chan *Task0) {
	for task := range done {
		// 根据实际需求处理已完成的任务
		fmt.Printf("处理完第 %d个任务了已经\n", task.ID)
	}
}
