package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func frontOrderNewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}
}

//func main() {
//	n1 := frontOrderNewTreeNode(1)
//	n2 := frontOrderNewTreeNode(2)
//	n3 := frontOrderNewTreeNode(3)
//	n4 := frontOrderNewTreeNode(4)
//	n5 := frontOrderNewTreeNode(5)
//	n6 := frontOrderNewTreeNode(6)
//	n7 := frontOrderNewTreeNode(7)
//	n1.Left = n2
//	n1.Right = n3
//	n2.Left = n4
//	n2.Right = n5
//	n3.Left = n6
//	n3.Right = n7
//	fmt.Println(test(n1))
//}
//
//func test(root *TreeNode) {
//	var stack []int
//	var Order func(node *TreeNode)
//	Order = func(node *TreeNode) {
//		if node == nil {
//			return
//		}
//		stack = append(stack, node.Val)
//		Order(node.Left)
//		Order(node.Right)
//	}
//}
