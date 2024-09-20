package main

import "sync"

// 节点
type bItem struct {
	Key int64
	Val interface{} // BPItem用于数据记录
}

// 节点
type bNode struct {
	MaxKey int64    //用于存储子树的最大关键字
	Nodes  []*bNode //节点的子树
	Items  []bItem  //叶子结点的数据记录
	Next   *bNode   //叶子结点中指向下一个叶子结点
}

type bTree struct {
	mutex sync.Mutex
	root  *bNode //指向B+树的根结点
	width int    //b+树的阶
	halfw int
}

func NewBTree(width int) *bTree {
	if width < 3 {
		width = 3
	}
	var b = &bTree{}
	b.root = NewLeafNode(width)
	b.width = width
	b.halfw = (width + 1) / 2
	return b
}

// 申请width+1是因为插入时可能暂时出现节点key大于申请width的情况,待后期再分裂处理
func NewLeafNode(width int) *bNode {
	var node = &bNode{}
	node.Items = make([]bItem, width+1)
	node.Items = node.Items[0:0]
	return node
}
