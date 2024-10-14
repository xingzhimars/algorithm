package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hNext := head.Next
	head.Next = swapPairs(hNext.Next)
	hNext.Next = head
	return hNext
}
