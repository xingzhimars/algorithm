package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head

	p1, p2 := dummyNode, dummyNode

	for i := 0; i < n; i++ {
		p1 = p1.Next
		if p1 == nil {
			return dummyNode.Next
		}
	}
	for p1 != nil {
		if p1.Next == nil {
			p2.Next = p2.Next.Next
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return dummyNode.Next
}
