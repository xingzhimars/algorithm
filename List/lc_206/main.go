package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, pre *ListNode
	cur := head

	for cur != nil {
		tmp := cur.Next
		if tmp == nil {
			newHead = cur
		}
		cur.Next = pre
		pre = cur
		// 这里不能cur = cur.Next，因为此时的cur.Next是pre，成环了
		cur = tmp
	}
	return newHead
}
