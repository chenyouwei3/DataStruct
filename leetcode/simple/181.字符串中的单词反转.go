package main

import "fmt"

func main() {
	message := "the sky is blue"
	fmt.Println(reverseMessage1(message))
}

func reverseMessage(message string) string {
	tempSlice := make([][]byte, len(message))
	for i := range tempSlice {
		tempSlice[i] = make([]byte, 0)
	}
	var j int
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			j++
		}
		if message[i] != ' ' {
			tempSlice[j] = append(tempSlice[j], message[i])
		}
	}
	var res string
	for _, s := range tempSlice {
		if string(s) == "" {
			continue
		}
		res += " " + string(s)
	}
	return res[1:]
}

func reverseMessage1(message string) string {
	if len(message) == 0 {
		return ""
	}
	if len(message) == 1 {
		return message
	}
	var tempSlice [][]byte
	var word []byte
	var temp byte
	for i := 0; i < len(message); i++ {
		if temp == ' ' {
			continue
		}
		if message[i] == ' ' {
			tempSlice = append(tempSlice, word)
			word = nil
			temp = message[i]
		} else {
			word = append(word, message[i])
			temp = message[i]
		}
	}
	// 添加最后一个单词到 tempSlice 中
	tempSlice = append(tempSlice, word)
	//反转
	reverse := make([]string, len(tempSlice))
	for i, v := range tempSlice {
		reverse[len(tempSlice)-1-i] = string(v)
	}
	var res string
	for _, s := range reverse {
		if s == "" {
			continue
		}
		res += " " + s
	}
	return res[1:]
}
