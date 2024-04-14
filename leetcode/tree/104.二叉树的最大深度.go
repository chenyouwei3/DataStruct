package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/* 创造新节点 */
func maxDepthNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := maxDepthNewTreeNode(1)
	n2 := maxDepthNewTreeNode(2)
	n3 := maxDepthNewTreeNode(3)
	n4 := maxDepthNewTreeNode(4)
	n5 := maxDepthNewTreeNode(5)
	n6 := maxDepthNewTreeNode(6)
	n7 := maxDepthNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println(maxDepth(n1))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepthMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
