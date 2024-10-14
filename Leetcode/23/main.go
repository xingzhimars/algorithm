package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, start, end int) *ListNode {
	if start > end {
		return nil
	}
	if start == end {
		return lists[start]
	}
	mid := start + (end-start)/2
	p1 := merge(lists, start, mid)
	p2 := merge(lists, mid+1, end)
	return mergeTwoList(p1, p2)
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	cur := dummyNode
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			cur.Next = p1
			p1 = p1.Next
		} else {
			cur.Next = p2
			p2 = p2.Next
		}
		cur = cur.Next
	}
	if p1 != nil {
		cur.Next = p1
	}
	if p2 != nil {
		cur.Next = p2
	}
	return dummyNode.Next
}
