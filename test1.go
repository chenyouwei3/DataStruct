package main

<<<<<<< HEAD
import "fmt"
=======
import (
	"fmt"
)
>>>>>>> 3b1c537 (9/20)

func main() {
	// 定义一个切片
	slice := []int{1, 2, 3, 4, 5}
<<<<<<< HEAD

	// 要删除的元素的索引
	index := 2

	// 使用切片的切片方式删除元素
	slice = append(slice[:index], slice[index+1:]...)

	fmt.Println(slice) // [1 2 4 5]
=======
	//element := 10
	// 要删除的元素的索引
	index := 2
	//insert := []int{element}
	// 使用切片的切片方式删除元素
	slice = append(slice[:index], slice[index+1:]...)
	//使用切片添加元素
	//mykou := append(slice[:index], append(insert, slice[index:]...)...)
	fmt.Println(slice) // [1 2 4 5]
	//fmt.Println(mykou)
>>>>>>> 3b1c537 (9/20)
}
