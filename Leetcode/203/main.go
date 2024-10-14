package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head

	pre := dummyNode
	cur := head
	for cur != nil {
		if cur.Val != val {
			pre = cur
			cur = cur.Next
		} else {
			pre.Next = cur.Next
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
