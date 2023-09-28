package main

import "fmt"

func main() {
	number := 10
	func(number int) {
		fmt.Println(number)
	}(number)

}
