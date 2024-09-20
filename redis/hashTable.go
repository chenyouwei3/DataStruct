package main

import "fmt"

type pair struct {
	key   int
	value string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{buckets: make([]*pair, 100)}
}

// 哈希函数
func (a *arrayHashMap) hashFunc(key int) int {
	return key % 100
}

// 查找
func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	pair := a.buckets[index]
	if pair == nil {
		return "not found"
	}
	return pair.value
}

// 增加
func (a *arrayHashMap) add(key int, value string) {
	pair := &pair{
		key:   key,
		value: value,
	}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

// 删除
func (a *arrayHashMap) deleted(key int) {
	index := a.hashFunc(key)
	a.buckets[index] = nil
}

// 获取所有键值对
func (a *arrayHashMap) getPairs() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

// 获取所有key
func (a *arrayHashMap) getKeys() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

// 获取所有value
func (a *arrayHashMap) getValues() []string {
	var strings []string
	for _, pair := range a.buckets {
		if pair != nil {
			strings = append(strings, pair.value)
		}
	}
	return strings
}

/*打印哈希表*/
func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.value)
		}
	}
}

func main() {
	newHashMap := newArrayHashMap()
	newHashMap.add(001, "cyw")
	newHashMap.add(002, "cx")
	newHashMap.add(003, "gdy")
	fmt.Println(newHashMap.get(001))
	fmt.Println(newHashMap.getValues())
	newHashMap.print()
}
