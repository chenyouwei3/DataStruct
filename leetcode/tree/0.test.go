package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 创造新节点 */
func testNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	n1 := frontOrderNewTreeNode(1)
	n2 := frontOrderNewTreeNode(2)
	n3 := frontOrderNewTreeNode(3)
	n4 := frontOrderNewTreeNode(4)
	n5 := frontOrderNewTreeNode(5)
	n6 := frontOrderNewTreeNode(6)
	n7 := frontOrderNewTreeNode(7)
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n3.Left = n6
	n3.Right = n7
	fmt.Println()
}

func test(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		temp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			temp = append(temp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, temp)
	}
	return res
}
