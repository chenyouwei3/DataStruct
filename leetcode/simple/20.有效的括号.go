package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println(s[:2], s[2:])
	isValid(s)
}

func isValid(s string) {
	var temp int32
	for i, v := range s {
		fmt.Println(v, temp, i)
		temp = v
	}
}
