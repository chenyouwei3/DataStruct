package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSymmetricNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := isSymmetricNewTreeNode(1)
	n2 := isSymmetricNewTreeNode(2)
	n3 := isSymmetricNewTreeNode(3)
	n4 := isSymmetricNewTreeNode(4)
	n5 := isSymmetricNewTreeNode(5)
	n6 := isSymmetricNewTreeNode(6)
	n7 := isSymmetricNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println(isSymmetric(n1))
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricMirrorTree(root.Left, root.Right)
}

func isSymmetricMirrorTree(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricMirrorTree(left.Left, right.Right) && isSymmetricMirrorTree(left.Right, right.Left)
}
