package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{
		Val: -110,
	}
	dummyNode.Next = head

	pre := dummyNode
	cur := head

	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		} else {
			pre.Next = cur
			pre = cur
			cur = cur.Next
		}
	}
	return dummyNode.Next
}
