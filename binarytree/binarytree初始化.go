package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode /* 创造新节点 */
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func leveOrder(root *TreeNode) []any {
	queue:=
}

func main() {
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	/* 插入节点 */
	p := NewTreeNode(0)
	n1.Left = p
	p.Left = n2
	/* 删除节点p */
	n1.Left = n2
}
