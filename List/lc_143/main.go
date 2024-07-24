package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	mid := getMid(head)
	p2 := reverse(mid.Next)
	mid.Next = nil

	p1 := head

	for p1 != nil && p2 != nil {
		n1 := p1.Next
		n2 := p2.Next

		p1.Next = p2
		p2.Next = n1

		p1 = n1
		p2 = n2
	}

}

func reverse(head *ListNode) *ListNode {
	var newHead, pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		if next == nil {
			newHead = cur
		}
		cur.Next = pre
		pre = cur
		cur = next
	}
	return newHead
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
