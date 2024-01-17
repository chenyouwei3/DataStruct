package main

/*hash冲突 */
/*开放寻址*/

type pair struct {
	key int
	val string
}

/*链式地址哈希表*/
type hashMapChain1 struct {
	pairSize    int      //键值对大小
	capacity    int      //哈希桶大小
	loadThree   float64  //触发扩容的负载因子阈值
	extendRatio int      //扩容倍数
	buckets     [][]pair //桶组数
	removed     pair     //删除标记
}

/*初始化哈希表*/
func newHashMapChain1() *hashMapChain1 {
	buckets := make([][]pair, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = make([]pair, 0)
	}
	return &hashMapChain1{
		pairSize:    0,       //键值对数量
		capacity:    4,       //哈希桶数量
		loadThree:   2 / 3.0, //负载因子阈值
		extendRatio: 2,       //扩容倍数
		buckets:     buckets, //桶数量
	}
}

/*哈希函数*/
func (m *hashMapChain1) hashFunc(key int) int {
	return key % m.capacity
}

/*负载因子*/
func (m *hashMapChain1) loadFactor() float64 {
	return float64(m.pairSize / m.capacity)
}

/*查询操作*/
func (m *hashMapChain1) get(key int) string {
	index := m.hashFunc(key)
	return ""
}
