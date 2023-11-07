package main

import (
	"fmt"
	"sync"
)

func worker1(link chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range link {
		fmt.Println(url)
	}
}

func main() {
	var yourLinksSlice = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	linkChan := make(chan string)
	wg := new(sync.WaitGroup)
	//开始任务
	for i := 0; i < 250; i++ {
		wg.Add(1)
		go worker1(linkChan, wg)
	}

	for _, link := range yourLinksSlice {
		linkChan <- link
	}
	close(linkChan)
	wg.Wait()
}
