package main

import "fmt"

//设计一个简单的hash表

// 键值对(定义demo)
type pair struct {
	key int
	val string
}

// 基于数组简易实现的哈希表
type arrayHashMap struct {
	buckets []*pair
}

// 初始化哈希表
func newArrayHashMap() *arrayHashMap {
	//初始化桶,自定义数量(200)
	buckets := make([]*pair, 200)
	return &arrayHashMap{buckets: buckets}
}

// 哈希函数(辅助找到桶)
func (a *arrayHashMap) hashFunction(key int) int {
	index := key % 200
	return index
}

// 查询
func (a *arrayHashMap) get(key int) string {
	index := a.hashFunction(key)
	pair := a.buckets[index]
	if pair == nil {
		return "NOT FOUND"
	}
	return pair.val
}

// 添加
func (a *arrayHashMap) put(key int, val string) {
	pair := &pair{key: key, val: val}
	index := a.hashFunction(key)
	a.buckets[index] = pair
}

// 删除
func (a *arrayHashMap) remove(key int) {
	index := a.hashFunction(key)
	//重置为空,代表删除
	a.buckets[index] = nil
}

// 获取所有键
func (a *arrayHashMap) getKey() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

// 获取所有值
func (a *arrayHashMap) getVal() []int {
	var values []int
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.key)
		}
	}
	return values
}

// 获得所有键值对
func (a *arrayHashMap) getPair() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

/* 打印哈希表 */
func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}
