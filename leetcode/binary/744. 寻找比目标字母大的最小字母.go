package main

import (
	"fmt"
)

func main() {
	letters := []byte{'c', 'f', 'j'}
	fmt.Println(letters)
	fmt.Println(nextGreatestLetter(letters, 'a'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	var res []byte
	for _, v := range letters {
		if v > target {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		return letters[0]
	}
	return res[0]
}
