package main

import (
	"crypto/md5"
	"fmt"
	"sort"
)

type HashRing struct {
	nodes    []int
	hashMap  map[int]string
	replicas int // 虚拟节点的数量
}

// NewHashRing 创建一个新的哈希环
func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		hashMap:  make(map[int]string),
		replicas: replicas,
	}
}

// AddNode 添加节点到哈希环
func (hr *HashRing) AddNode(node string) {
	for i := 0; i < hr.replicas; i++ {
		hash := hr.hash(node + fmt.Sprintf("%d", i))
		hr.hashMap[hash] = node
		hr.nodes = append(hr.nodes, hash)
	}
	sort.Ints(hr.nodes)
}

// RemoveNode 从哈希环中删除节点
func (hr *HashRing) RemoveNode(node string) {
	for i := 0; i < hr.replicas; i++ {
		hash := hr.hash(node + fmt.Sprintf("%d", i))
		delete(hr.hashMap, hash)
	}
	hr.updateNodes()
}

// updateNodes 更新节点列表
func (hr *HashRing) updateNodes() {
	hr.nodes = []int{}
	for k := range hr.hashMap {
		hr.nodes = append(hr.nodes, k)
	}
	sort.Ints(hr.nodes)
}

// GetNode 获取对应key的节点
func (hr *HashRing) GetNode(key string) string {
	if len(hr.nodes) == 0 {
		return ""
	}
	hash := hr.hash(key)
	idx := sort.Search(len(hr.nodes), func(i int) bool {
		return hr.nodes[i] >= hash
	})
	if idx == len(hr.nodes) {
		idx = 0 // 如果找到的索引超出范围，返回第一个节点
	}
	return hr.hashMap[hr.nodes[idx]]
}

// hash 计算哈希值
//func (hr *HashRing) hash(key string) int {
//	h := md5.Sum([]byte(key))
//	// 使用更多字节并取模以避免溢出
//	return int(h[0]) + int(h[1])<<8 + int(h[2])<<16 + int(h[3])<<24 + int(h[4])<<32 + int(h[5])<<40 + int(h[6])<<48 + int(h[7])<<56
//}

// hash 计算哈希值
func (hr *HashRing) hash(key string) int {
	h := md5.Sum([]byte(key))
	hashInt := 0
	for _, b := range h {
		hashInt = hashInt*31 + int(b) // 使用31作为乘子，以减少碰撞
	}
	return hashInt
}

func main() {
	ring := NewHashRing(100) // 使用 100 个虚拟节点
	ring.AddNode("Node1")
	ring.AddNode("Node2")
	ring.AddNode("Node3")
	ring.AddNode("Node4")
	ring.AddNode("Node5")

	keys := []string{"key1", "key2", "key3", "key4", "key5", "key6", "key7", "key8", "key9", "key10"}
	for _, key := range keys {
		node := ring.GetNode(key)
		fmt.Printf("Key %s is mapped to node %s\n", key, node)
	}
}
