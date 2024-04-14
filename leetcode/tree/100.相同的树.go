package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSameTreeNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := isSameTreeNewTreeNode(1)
	n2 := isSameTreeNewTreeNode(2)
	n3 := isSameTreeNewTreeNode(3)
	n4 := isSameTreeNewTreeNode(4)
	n5 := isSameTreeNewTreeNode(5)
	n6 := isSameTreeNewTreeNode(6)
	n7 := isSameTreeNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println(isSameTree1(n1, n1))
}

func isSameTree1(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree1(p.Right, q.Right) && isSameTree1(p.Left, q.Left)
	}
	return false
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	return p.Val == q.Val && isSameTree2(p.Left, q.Left) && isSameTree2(p.Right, q.Right)
}
