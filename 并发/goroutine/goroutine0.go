package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}

func main() {
	go say("Go")                // 以协程方式执行say函数
	say("noGo")                 // 以普通方式执行say函数
	time.Sleep(5 * time.Second) // 睡眠5秒：防止协程未执行完毕，主程序退出
}
