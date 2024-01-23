package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func inorderNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := inorderNewTreeNode(1)
	n2 := inorderNewTreeNode(2)
	n3 := inorderNewTreeNode(3)
	n4 := inorderNewTreeNode(4)
	n5 := inorderNewTreeNode(5)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	fmt.Println(inorderTraversal(n1))
}

// 访问优先级:左子树->跟节点->右子树
func inorderTraversal(root *TreeNode) []int {
	var stack []int
	var inorderFunc func(node *TreeNode)
	inorderFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorderFunc(node.Left)
		stack = append(stack, node.Val)
		inorderFunc(node.Right)
	}
	inorderFunc(root)
	return stack
}

func inorderTraversal1(root *TreeNode) (stack []int) {

	var inorderFunc func(node *TreeNode)
	inorderFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorderFunc(node.Left)
		stack = append(stack, node.Val)
		inorderFunc(node.Right)
	}
	inorderFunc(root)
	return stack
}
