package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/* 创造新节点 */
func minDepthNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := minDepthNewTreeNode(1)
	n2 := minDepthNewTreeNode(2)
	n3 := minDepthNewTreeNode(3)
	n4 := minDepthNewTreeNode(4)
	n5 := minDepthNewTreeNode(5)
	n6 := minDepthNewTreeNode(6)
	n7 := minDepthNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println(minDepth(n1))
}
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	return minDepthMax(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minDepthMax(a, b int) int {
	if a > b {
		return b
	}
	return a
}
