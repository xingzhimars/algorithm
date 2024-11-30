package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	cur := dummy

	for i := 0; i < left; i++ {
		pre = cur
		cur = cur.Next
	}

	for i := 0; i < right-left; i++ {
		nx := cur.Next
		cur.Next = nx.Next
		// 头插法
		nx.Next = pre.Next
		pre.Next = nx
	}
	return dummy.Next
}
