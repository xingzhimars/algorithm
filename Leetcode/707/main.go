package main

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

// 错误
type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	list := MyLinkedList{
		head: new(Node),
		tail: new(Node),
		size: 0,
	}
	list.head.Next = list.tail
	list.tail.Prev = list.head
	return list
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		Val: val,
	}
	next := this.head.Next

	newNode.Next = next
	next.Prev = newNode

	this.head.Next = newNode
	newNode.Prev = this.head

}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		Val: val,
	}

	lastNode := this.tail.Prev

	newNode.Next = this.tail
	this.tail.Prev = newNode

	lastNode.Next = newNode
	newNode.Prev = lastNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	cur := this.head.Next
	for cur != this.tail && index > 0 {
		cur = cur.Next
		index--
	}
	if index > 0 {
		return
	}
	newNode := &Node{
		Val: val,
	}
	cur.Prev.Next = newNode
	newNode.Prev = cur.Prev

	newNode.Next = cur
	cur.Prev = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head.Next
	for cur != this.tail && index > 0 {
		cur = cur.Next
		index--
	}
	if index > 0 {
		return
	}
	cur.Prev.Next = cur.Next
	if cur.Next != nil {
		cur.Next.Prev = cur.Prev
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
