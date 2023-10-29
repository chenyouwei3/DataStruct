package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	//写入
	m.Store("demo1", 18)
	m.Store("demo2", 20)
	//读取
	age, _ := m.Load("demo1")
	fmt.Println(age)
	//遍历
	m.Range(func(key, value interface{}) bool {
		name := key.(string)
		age := value.(int)
		fmt.Println(name, age)
		return true
	})
	//删除
	m.Delete("demo2")
	age, ok := m.Load("demo2")
	fmt.Println(age, ok)
	//读取或者写入
	m.LoadOrStore("demo3", 100)
	age, _ = m.Load("demo3")
	fmt.Println(age)
}
