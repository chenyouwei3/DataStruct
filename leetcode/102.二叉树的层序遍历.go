package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 创造新节点 */
func levelOrderNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := levelOrderNewTreeNode(1)
	n2 := levelOrderNewTreeNode(2)
	n3 := levelOrderNewTreeNode(3)
	n4 := levelOrderNewTreeNode(4)
	n5 := levelOrderNewTreeNode(5)
	n6 := levelOrderNewTreeNode(6)
	n7 := levelOrderNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	//fmt.Println(levelOrder(n1)) //1,2,3,4,5
	//frontOrder(n1) //1,2,4,5,3
	//inOrder(n1) //4 2 5 1 3
	//backOrder(n1) //4 5 2 3 1
	fmt.Println(levelOrder(n1))
}

func levelOrder(root *TreeNode) [][]int {

}
