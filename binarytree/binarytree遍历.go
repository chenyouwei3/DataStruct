package main

import (
	"container/list"
	"fmt"
)

var binaryTreeResSlice []int

//	type TreeNode struct {
//		Val   int
//		Left  *TreeNode
//		Right *TreeNode
//	}
//
//	func newTreeNode(v int) *TreeNode {
//		return &TreeNode{
//			Val:   v,
//			Left:  nil,
//			Right: nil,
//		}
//	}
func main() {
	n1 := newTreeNode(1)
	n2 := newTreeNode(2)
	n3 := newTreeNode(3)
	n4 := newTreeNode(4)
	n5 := newTreeNode(5)
	n6 := newTreeNode(6)
	n7 := newTreeNode(7)
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

/*层序遍历*/
func levelOrder(root *TreeNode) []int {
	queue := list.New() //初始化队列  加入根节点
	queue.PushBack(root)
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*TreeNode)
		binaryTreeResSlice = append(binaryTreeResSlice, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return binaryTreeResSlice
}

/*深度优先遍历*/
/*前序遍历*/
func frontOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级:根节点->左子树->右子树
	binaryTreeResSlice = append(binaryTreeResSlice, node.Val)
	fmt.Println(binaryTreeResSlice)
	frontOrder(node.Left)
	frontOrder(node.Right)
}

/*中序遍历*/
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	//访问优先级:左子树->跟节点->右子树
	inOrder(node.Left)
	binaryTreeResSlice = append(binaryTreeResSlice, node.Val)
	inOrder(node.Right)
}

/*后序遍历*/
func backOrder(node *TreeNode) {
	if node == nil {
		return
	}
	//访问优先级:左子树->右子树->根节点
	backOrder(node.Left)
	backOrder(node.Right)
	binaryTreeResSlice = append(binaryTreeResSlice, node.Val)
}
