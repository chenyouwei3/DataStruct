package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}
func main() {
	n0 := newTreeNode(8)
	n1 := newTreeNode(4)
	n2 := newTreeNode(12)
	n0.Left = n1
	n0.Right = n2
	n3 := newTreeNode(2)
	n4 := newTreeNode(6)
	n1.Left = n3
	n1.Right = n4
	n5 := newTreeNode(10)
	n6 := newTreeNode(14)
	n2.Left = n5
	n2.Right = n6
	n7 := newTreeNode(1)
	n8 := newTreeNode(3)
	n3.Left = n7
	n3.Right = n8
	n9 := newTreeNode(5)
	n10 := newTreeNode(7)
	n4.Left = n9
	n4.Right = n10
	n11 := newTreeNode(9)
	n12 := newTreeNode(11)
	n5.Left = n11
	n5.Right = n12
	n13 := newTreeNode(13)
	n14 := newTreeNode(15)
	n6.Left = n13
	n6.Right = n14

}
