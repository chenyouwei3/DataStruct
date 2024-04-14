package main

//
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	buckets := make(map[*ListNode]int)
	for cur1 := headA; cur1 != nil; cur1 = cur1.Next {
		buckets[cur1]++
	}
	for cur2 := headB; cur2 != nil; cur2 = cur2.Next {
		if _, ok := buckets[cur2]; ok {
			return cur2
		}
	}
	return nil
}
