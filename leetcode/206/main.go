package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		if tmp == nil {
			newHead = cur
		}
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return newHead
}
