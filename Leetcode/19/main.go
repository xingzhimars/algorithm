package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	p1, p2 := dummy, dummy

	for i := 0; i < n; i++ {
		p1 = p1.Next
		if p1 == nil {
			return dummy.Next
		}
	}

	for p1 != nil {
		if p1.Next == nil {
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	if p2.Next != nil {
		p2.Next = p2.Next.Next
	}

	return dummy.Next
}
