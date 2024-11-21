package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这道题目是有重复的元素，就不能选，而不是选一个
// 那就是说需要判断 cur.Val 是否等于 cur.Next.Val，不等的话，就可以选cur
func deleteDuplicatesErr(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val: -110,
	}
	dummy.Next = head

	pre := dummy
	cur := head

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		cur = cur.Next

		pre.Next = cur
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val: -110,
	}
	dummy.Next = head

	pre := dummy
	cur := head

	for cur != nil {
		if cur.Next == nil || cur.Val != cur.Next.Val {
			pre.Next = cur
			pre = cur
		}
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		cur = cur.Next
	}
	pre.Next = nil
	return dummy.Next
}
