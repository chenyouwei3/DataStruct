package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/* 创造新节点 */
func mirrorTreeNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := mirrorTreeNewTreeNode(1)
	n2 := mirrorTreeNewTreeNode(2)
	n3 := mirrorTreeNewTreeNode(3)
	n4 := mirrorTreeNewTreeNode(4)
	n5 := mirrorTreeNewTreeNode(5)
	n6 := mirrorTreeNewTreeNode(6)
	n7 := mirrorTreeNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println(mirrorTree(n1))
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left //交换
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}
