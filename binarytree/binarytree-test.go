package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func newTreeNode(v int) *TreeNode {
//	return &TreeNode{
//		Val:   v,
//		Left:  nil,
//		Right: nil,
//	}
//}

func main() {
	n1 := newTreeNode(1)
	n2 := newTreeNode(2)
	n3 := newTreeNode(3)
	n4 := newTreeNode(4)
	n5 := newTreeNode(5)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
}
