package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	h2 := mid.Next
	mid.Next = nil

	l1 := sortList(head)
	l2 := sortList(h2)
	return merge(l1, l2)
}

func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}

func merge(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	dummy := new(ListNode)
	cur := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			cur.Next = h1
			h1 = h1.Next
		} else {
			cur.Next = h2
			h2 = h2.Next
		}
		cur = cur.Next
	}
	if h1 != nil {
		cur.Next = h1
	}
	if h2 != nil {
		cur.Next = h2
	}
	return dummy.Next
}
