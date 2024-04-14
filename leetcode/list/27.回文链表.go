package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l3 := new(ListNode)
	l3.Val = 4
	l4 := new(ListNode)
	l4.Val = 3
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4

}

func isPalindrome(head *ListNode) bool {
	var temp []int
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		if temp[i] != temp[j] {
			return false // 如果发现不匹配的值，则返回 false
		}
	}
	return true
}
