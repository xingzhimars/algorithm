package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val: 110,
	}
	dummy.Next = head

	tail := dummy
	for tail.Next != nil {
		p1 := tail.Next.Next
		for p1 != nil && p1.Val == tail.Next.Val {
			p1 = p1.Next
		}
		if tail.Next.Next == p1 {
			// 中间只有一个值
			tail = tail.Next
		} else {
			tail.Next = p1
		}
	}
	return dummy.Next
}
