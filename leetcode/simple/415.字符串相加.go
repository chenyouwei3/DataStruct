package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	str1 := "3876620623801494171"
	str2 := "6529364523802684779"
	fmt.Println(addStrings(str1, str2))
	//fmt.Println(addStringsTest(2))
}

func addStrings(num1 string, num2 string) string {
	var total1, total2 int
	for i, str1 := range num1 {
		total1 += int(math.Pow10(len(num1)-1-i)) * int(str1-'0')
	}
	for i, str2 := range num2 {
		total2 += int(math.Pow10(len(num2)-1-i)) * int(str2-'0')
	}
	sum := total1 + total2
	if sum < 0 {
		return string(math.MaxInt64)
	}
	return strconv.Itoa(total2 + total1)
}

//func addStringsTest(exponent int32) int32 {
//	result := 1
//	for i := 0; i < exponent; i++ {
//		result *= 10
//	}
//	return result
//}
