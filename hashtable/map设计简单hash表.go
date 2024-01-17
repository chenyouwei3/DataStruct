package main

import "fmt"

// 设计一个简单的hash表

//type pair struct {
//	key int
//	val string
//}

/*基于数组简易实现的哈希表*/
type arrayHashMap struct {
	buckets []*pair
}

/*初始化哈希表*/
func newArrayHashMap() *arrayHashMap {
	//初始化桶,包含多少个桶
	buckets := make([]*pair, 100)
	return &arrayHashMap{buckets: buckets}
}

/* 哈希函数*/
func (a *arrayHashMap) hashFunc(key int) int {
	//index := key % 100 //限制桶的数量
	return key % 100
}

/*查询操作*/
func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "not found"
	}
	return pair.val
}

/*添加操作*/
func (a *arrayHashMap) add(key int, value string) {
	pair := &pair{
		key: key,
		val: value,
	}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

/*删除操作*/
func (a *arrayHashMap) deleted(key int) {
	index := a.hashFunc(key)
	//重置为空相当于删除
	a.buckets[index] = nil
}

/*获取所有键值对*/
func (a *arrayHashMap) getPairs() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

/*获取所有key*/
func (a *arrayHashMap) getKey() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

/*获取所有value*/
func (a *arrayHashMap) getValue() []string {
	var strings []string
	for _, pair := range a.buckets {
		if pair != nil {
			strings = append(strings, pair.val)
		}
	}
	return strings
}

/*打印哈希表*/
func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}

func main() {
	newHashMap := newArrayHashMap()
	newHashMap.add(001, "cyw")
	newHashMap.add(002, "cx")
	newHashMap.add(003, "gdy")
	fmt.Println(newHashMap.get(001))
	fmt.Println(newHashMap.getValue())
	newHashMap.print()
}
