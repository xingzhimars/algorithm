package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head

	// tail插入的尾节点
	tail, cur := head, head.Next
	for cur != nil {
		if tail.Val <= cur.Val {
			tail = tail.Next
			cur = tail.Next
		} else {
			pre := dummy
			for pre.Next.Val <= cur.Val {
				pre = pre.Next
			}
			// 插入在pre后面，同时需要删除cur节点
			tail.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
			cur = tail.Next
		}
	}
	return dummy.Next
}
