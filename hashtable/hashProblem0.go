package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*hash冲突 */
/*链式地址*/

type pair struct {
	key int
	val string
}

// 设计一个简单的hash表
type hashMapChain struct {
	pairSize    int      //键值对数量
	capacity    int      //哈希桶数量
	loadThree   float64  //触发扩容的负载因子阈值
	extendRatio int      //扩容倍数
	buckets     [][]pair //桶组数
}

/*初始化哈希表*/
func newHashMapChain() *hashMapChain {
	buckets := make([][]pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair, 0)
	}
	return &hashMapChain{
		pairSize:    0,
		capacity:    4,
		loadThree:   2 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
	}
}

/*哈希函数*/
func (m *hashMapChain) hashFunc(key int) int {
	return key % m.capacity
}

/*负载因子*/
func (m *hashMapChain) loadFactor() float64 {
	return float64(m.pairSize / m.capacity)
}

/*查询操作*/
func (m *hashMapChain) get(key int) string {
	index := m.hashFunc(key)
	bucket := m.buckets[index]
	//遍历桶,找到key返回对应的val
	for _, pair := range bucket {
		if pair.key == key {
			return pair.val
		}
	}
	return ""
}

/*添加操作*/
func (m *hashMapChain) add(key int, value string) {
	//检查负载因子是否超过阈值
	if m.loadFactor() > m.loadThree {
		m.extend()
	}
	index := m.hashFunc(key)
	for _, pair := range m.buckets[index] {
		if pair.key == key {
			pair.val = value
			return
		}
	}
	pair := pair{
		key: key,
		val: value,
	}
	m.buckets[index] = append(m.buckets[index], pair)
	m.pairSize += 1

}

/*删除操作*/
func (m *hashMapChain) deleted(key int) {
	index := m.hashFunc(key)
	//遍历桶,从中删除键值对
	for i, pair := range m.buckets[index] {
		if pair.key == key {
			//删除桶的键值对
			m.buckets[index] = append(m.buckets[index][:i], m.buckets[index][i+1:]...)
			m.pairSize -= 1
			break
		}
	}
}

/*扩容哈希表*/
func (m *hashMapChain) extend() {
	//暂存原始哈希表
	tempBuckets := make([][]pair, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		tempBuckets[i] = make([]pair, len(m.buckets[i]))
		copy(tempBuckets[i], m.buckets[i])
	}
	m.capacity *= m.extendRatio
	m.buckets = make([][]pair, m.capacity)
	for i := 0; i < m.capacity; i++ {
		m.buckets[i] = make([]pair, 0)
	}
	m.pairSize = 0
	// 将键值对从原哈希表搬运至新哈希表
	for _, bucket := range tempBuckets {
		for _, pair := range bucket {
			m.add(pair.key, pair.val)
		}
	}
}

/* 打印哈希表 */
func (m *hashMapChain) print() {
	var builder strings.Builder
	for _, bucket := range m.buckets {
		builder.WriteString("[")
		for _, p := range bucket {
			builder.WriteString(strconv.Itoa(p.key) + " -> " + p.val + " ")
		}
		builder.WriteString("]")
		fmt.Println(builder.String())
		builder.Reset()
	}
}
func main() {
	newHashMap := newHashMapChain()
	newHashMap.add(000, "cyw")
	newHashMap.add(001, "cx")
	newHashMap.add(002, "gdy")
	fmt.Println(newHashMap.get(000), "111")
	fmt.Println(newHashMap.get(001))
	fmt.Println(newHashMap.get(002))
	newHashMap.print()

}
