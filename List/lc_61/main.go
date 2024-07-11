package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	k = k % count

	if k == 0 {
		return head
	}

	var newHead *ListNode
	cur = head
	for i := 0; i < count-k-1; i++ {
		cur = cur.Next
	}
	newHead = cur.Next
	cur.Next = nil

	cur = newHead
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	if cur != nil {
		cur.Next = head
	}

	return newHead

}
