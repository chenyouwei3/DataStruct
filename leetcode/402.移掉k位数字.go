package main

import "fmt"

type Test struct {
	Name string
	why  string
}

type newTest struct {
	nickname string
	test     Test
}

func main() {
	test1 := &Test{
		Name: "111",
		why:  "22",
	}
	test2 := &newTest{
		nickname: "333",
		test:     *test1,
	}
	fmt.Println(test2.test.why)
}

//func removeKdigits(num string, k int) string {
//	if
//}
