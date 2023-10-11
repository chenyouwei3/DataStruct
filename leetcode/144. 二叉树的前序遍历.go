package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := NewTreeNode(1)
	n2 := NewTreeNode(2)
	n3 := NewTreeNode(3)
	n4 := NewTreeNode(4)
	n5 := NewTreeNode(5)
	n6 := NewTreeNode(6)
	n7 := NewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	preorderTraversal(n1)
}

/* 递归 */
func preorderTraversal(root *TreeNode) []int {
	var stack []int
	var funcpreorder func(*TreeNode)
	funcpreorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		stack = append(stack, node.Val)
		funcpreorder(node.Left)
		funcpreorder(node.Right)
	}
	funcpreorder(root)
	fmt.Println(stack)
	return stack
}
