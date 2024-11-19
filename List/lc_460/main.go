package main

type Node struct {
	Key, Val, Freq int
	Prev, Next     *Node
}

type DoubleListNode struct {
	head, tail *Node
}

func NewDoubleListNode() *DoubleListNode {
	dl := &DoubleListNode{
		head: new(Node),
		tail: new(Node),
	}
	dl.head.Next = dl.tail
	dl.tail.Prev = dl.head
	return dl
}

func (dl *DoubleListNode) addFirst(node *Node) {
	node.Next = dl.head.Next
	dl.head.Next.Prev = node

	dl.head.Next = node
	node.Prev = dl.head
}

func (dl *DoubleListNode) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (dl *DoubleListNode) IsEmpty() bool {
	return dl.head.Next == dl.tail
}

type LFUCache struct {
	cache              map[int]*Node
	freq               map[int]*DoubleListNode
	cap, size, minFreq int
}

func Constructor(capacity int) LFUCache {
	lfu := LFUCache{
		cap:   capacity,
		cache: make(map[int]*Node),
		freq:  make(map[int]*DoubleListNode),
	}
	return lfu
}

func (lfu *LFUCache) Get(key int) int {
	node, ok := lfu.cache[key]
	if !ok {
		return -1
	}
	lfu.incFreq(node)
	return node.Val
}

func (lfu *LFUCache) Put(key int, value int) {
	node, ok := lfu.cache[key]
	if ok {
		node.Val = value
		lfu.incFreq(node)
		return
	}
	if lfu.size >= lfu.cap {
		dl := lfu.freq[lfu.minFreq]
		lastNode := dl.tail.Prev
		dl.remove(lastNode)
		delete(lfu.cache, lastNode.Key)
		lfu.size--
	}
	newNode := &Node{
		Key:  key,
		Val:  value,
		Freq: 1,
	}
	lfu.cache[key] = newNode
	if lfu.freq[1] == nil {
		lfu.freq[1] = NewDoubleListNode()
	}
	lfu.freq[1].addFirst(newNode)
	lfu.size++
	lfu.minFreq = 1
}

func (lfu *LFUCache) incFreq(node *Node) {
	nodeFreq := node.Freq
	lfu.freq[nodeFreq].remove(node)
	if lfu.minFreq == nodeFreq && lfu.freq[nodeFreq].IsEmpty() {
		delete(lfu.freq, nodeFreq)
		lfu.minFreq++
	}
	node.Freq++
	if lfu.freq[node.Freq] == nil {
		lfu.freq[node.Freq] = NewDoubleListNode()
	}
	lfu.freq[node.Freq].addFirst(node)
}
