package main

import (
	"fmt"
	"strconv"
)

/*hash冲突 */
/*开放寻址*/
//「开放寻址 open addressing」不引入额外的数据结构，而是通过“多次探测”来处理哈希冲突，探测方式主要包括线性探测、平方探测、多次哈希等。
//线性探测采用固定步长的线性搜索来进行探测，其操作方法与普通哈希表有所不同。
//插入元素：通过哈希函数计算数组索引，若发现桶内已有元素，则从冲突位置向后线性遍历（步长通常为 1 ）直至找到空位，将元素插入其中。
//查找元素：若发现哈希冲突，则使用相同步长向后线性遍历，直到找到对应元素，返回 value 即可；如果遇到空位，说明目标键值对不在哈希表中，返回 None 。
//不能直接删除元素。删除元素会在数组内产生一个空位，当查找该空位之后的元素时，该空位可能导致程序误判元素不存在。为此，通常需要借助一个标志位来标记已删除元素。
//容易产生聚集。数组内连续被占用位置越长，这些连续位置发生哈希冲突的可能性越大，进一步促使这一位置的聚堆生长，形成恶性循环，最终导致增删查改操作效率劣化。

type pair struct {
	key int
	val string
}

/* 链式地址哈希表 */
type hashMapOpenAddressing struct {
	size        int     // 键值对数量
	capacity    int     // 哈希表容量
	loadThres   float64 // 触发扩容的负载因子阈值
	extendRatio int     // 扩容倍数
	buckets     []pair  // 桶数组
	removed     pair    // 删除标记
}

/* 构造方法 */
func newHashMapOpenAddressing() *hashMapOpenAddressing {
	buckets := make([]pair, 4)
	return &hashMapOpenAddressing{
		size:        0,
		capacity:    4,
		loadThres:   2 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
		removed: pair{
			key: -1,
			val: "-1",
		},
	}
}

/* 哈希函数 */
func (m *hashMapOpenAddressing) hashFunc(key int) int {
	return key % m.capacity
}

/* 负载因子 */
func (m *hashMapOpenAddressing) loadFactor() float64 {
	return float64(m.size) / float64(m.capacity)
}

/* 查询操作 */
func (m *hashMapOpenAddressing) get(key int) string {
	idx := m.hashFunc(key)
	// 线性探测，从 index 开始向后遍历
	for i := 0; i < m.capacity; i++ {
		// 计算桶索引，越过尾部返回头部
		j := (idx + 1) % m.capacity
		// 若遇到空桶，说明无此 key ，则返回 null
		if m.buckets[j] == (pair{}) {
			return ""
		}
		// 若遇到指定 key ，则返回对应 val
		if m.buckets[j].key == key && m.buckets[j] != m.removed {
			return m.buckets[j].val
		}
	}
	// 若未找到 key 则返回空字符串
	return ""
}

/* 添加操作 */
func (m *hashMapOpenAddressing) put(key int, val string) {
	// 当负载因子超过阈值时，执行扩容
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	idx := m.hashFunc(key)
	// 线性探测，从 index 开始向后遍历
	for i := 0; i < m.capacity; i++ {
		// 计算桶索引，越过尾部返回头部
		j := (idx + i) % m.capacity
		// 若遇到空桶、或带有删除标记的桶，则将键值对放入该桶
		if m.buckets[j] == (pair{}) || m.buckets[j] == m.removed {
			m.buckets[j] = pair{
				key: key,
				val: val,
			}
			m.size += 1
			return
		}
		// 若遇到指定 key ，则更新对应 val
		if m.buckets[j].key == key {
			m.buckets[j].val = val
		}
	}
}

/* 删除操作 */
func (m *hashMapOpenAddressing) remove(key int) {
	idx := m.hashFunc(key)
	// 遍历桶，从中删除键值对
	// 线性探测，从 index 开始向后遍历
	for i := 0; i < m.capacity; i++ {
		// 计算桶索引，越过尾部返回头部
		j := (idx + 1) % m.capacity
		// 若遇到空桶，说明无此 key ，则直接返回
		if m.buckets[j] == (pair{}) {
			return
		}
		// 若遇到指定 key ，则标记删除并返回
		if m.buckets[j].key == key {
			m.buckets[j] = m.removed
			m.size -= 1
		}
	}
}

/* 扩容哈希表 */
func (m *hashMapOpenAddressing) extend() {
	// 暂存原哈希表
	tmpBuckets := make([]pair, len(m.buckets))
	copy(tmpBuckets, m.buckets)
	// 初始化扩容后的新哈希表
	m.capacity *= m.extendRatio
	m.buckets = make([]pair, m.capacity)
	m.size = 0
	// 将键值对从原哈希表搬运至新哈希表
	for _, p := range tmpBuckets {
		if p != (pair{}) && p != m.removed {
			m.put(p.key, p.val)
		}
	}
}

/* 打印哈希表 */
func (m *hashMapOpenAddressing) print() {
	for _, p := range m.buckets {
		if p != (pair{}) {
			fmt.Println(strconv.Itoa(p.key) + " -> " + p.val)
		} else {
			fmt.Println("nil")
		}
	}
}
