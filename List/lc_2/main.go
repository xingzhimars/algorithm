package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummyNode := &ListNode{}
	cur := dummyNode
	carry := 0

	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		carry += v1 + v2
		newNode := &ListNode{
			Val: carry % 10,
		}
		carry /= 10
		cur.Next = newNode
		cur = cur.Next
	}
	if carry > 0 {
		cur.Next = &ListNode{
			Val: carry,
		}
	}
	return dummyNode.Next
}
