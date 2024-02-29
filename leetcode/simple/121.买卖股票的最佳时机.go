package main

import "fmt"

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}

func maxProfit(prices []int) int {
	maxLR, minJG, res := 0, prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minJG {
			minJG = prices[i]
		}
		res = prices[i] - minJG
		if res > maxLR {
			maxLR = res
		}
	}
	return maxLR
}

func maxProfit1(prices []int) int {
	var res int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			res = maxmaxProfit1(prices[j]-prices[i], res)
		}
	}
	if res < 0 {
		return 0
	}
	return res
}

func maxmaxProfit1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
