package main

import (
	"fmt"
)

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func backOrderNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := backOrderNewTreeNode(1)
	n2 := backOrderNewTreeNode(2)
	n3 := backOrderNewTreeNode(3)
	n4 := backOrderNewTreeNode(4)
	n5 := backOrderNewTreeNode(5)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	fmt.Println(backOrderTraversal(n1))
}

func backOrderTraversal(root *TreeNode) (res []int) {
	var backOrder func(node *TreeNode)
	backOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		backOrder(node.Left)
		backOrder(node.Right)
		res = append(res, node.Val)
	}
	backOrder(root)
	return
}
