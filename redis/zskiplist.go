package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxLevel = 16  // 最大层数
	P        = 0.5 // 每层节点的概率因子
)

// SkipListNode 表示跳表中的节点
type SkipListNode struct {
	value  int              // 节点的值
	levels []*SkipListLevel // 节点在不同层的索引
}

// SkipListLevel 表示跳表中每一层的索引
type SkipListLevel struct {
	next *SkipListNode // 下一层节点
}

// SkipList 表示跳表结构
type SkipList struct {
	header *SkipListNode // 表头节点
	length int           // 当前跳表的长度
	height int           // 当前跳表的高度
}

// NewSkipList 创建一个新的跳表
func NewSkipList() *SkipList {
	// 使用最大层数初始化表头节点
	header := NewSkipListNode(MaxLevel, 0)
	return &SkipList{
		header: header,
		length: 0,
		height: 1,
	}
}

// NewSkipListNode 创建一个新的跳表节点
func NewSkipListNode(level int, value int) *SkipListNode {
	node := &SkipListNode{
		value:  value,
		levels: make([]*SkipListLevel, level),
	}
	for i := 0; i < level; i++ {
		node.levels[i] = &SkipListLevel{}
	}
	return node
}

// randomLevel 生成一个随机的层数
func randomLevel() int {
	level := 1
	for rand.Float64() < P && level < MaxLevel {
		level++
	}
	return level
}

// Insert 插入一个新的值到跳表中
func (sl *SkipList) Insert(value int) {
	update := make([]*SkipListNode, MaxLevel)
	curr := sl.header

	// 从顶层开始遍历找到插入位置
	for i := sl.height - 1; i >= 0; i-- {
		for curr.levels[i].next != nil && curr.levels[i].next.value < value {
			curr = curr.levels[i].next
		}
		update[i] = curr
	}

	// 生成新节点的层数
	newLevel := randomLevel()
	if newLevel > sl.height {
		for i := sl.height; i < newLevel; i++ {
			update[i] = sl.header
		}
		sl.height = newLevel
	}

	// 插入新节点
	newNode := NewSkipListNode(newLevel, value)
	for i := 0; i < newLevel; i++ {
		newNode.levels[i].next = update[i].levels[i].next
		update[i].levels[i].next = newNode
	}

	sl.length++
}

// Search 查找跳表中是否存在某个值
func (sl *SkipList) Search(value int) bool {
	curr := sl.header
	for i := sl.height - 1; i >= 0; i-- {
		for curr.levels[i].next != nil && curr.levels[i].next.value < value {
			curr = curr.levels[i].next
		}
	}
	curr = curr.levels[0].next
	return curr != nil && curr.value == value
}

// Delete 从跳表中删除一个值
func (sl *SkipList) Delete(value int) {
	update := make([]*SkipListNode, MaxLevel)
	curr := sl.header

	// 查找要删除的节点
	for i := sl.height - 1; i >= 0; i-- {
		for curr.levels[i].next != nil && curr.levels[i].next.value < value {
			curr = curr.levels[i].next
		}
		update[i] = curr
	}

	curr = curr.levels[0].next
	if curr != nil && curr.value == value {
		for i := 0; i < sl.height; i++ {
			if update[i].levels[i].next == curr {
				update[i].levels[i].next = curr.levels[i].next
			}
		}

		// 下降高度
		for sl.height > 1 && sl.header.levels[sl.height-1].next == nil {
			sl.height--
		}
		sl.length--
	}
}

// Print 打印跳表的内容
func (sl *SkipList) Print() {
	for i := sl.height - 1; i >= 0; i-- {
		fmt.Printf("Level %d: ", i)
		curr := sl.header.levels[i].next
		for curr != nil {
			fmt.Printf("%d ", curr.value)
			curr = curr.levels[i].next
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子
	sl := NewSkipList()

	// 插入一些值
	sl.Insert(3)
	sl.Insert(6)
	sl.Insert(7)
	sl.Insert(9)
	sl.Insert(12)
	sl.Insert(19)
	sl.Insert(17)

	// 打印跳表内容
	sl.Print()

	// 查找值
	fmt.Println("Search 9:", sl.Search(9))   // Output: true
	fmt.Println("Search 15:", sl.Search(15)) // Output: false

	// 删除值
	sl.Delete(19)
	sl.Print()
}
