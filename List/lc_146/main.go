package main

type LRUCache struct {
	head, tail *Node
	capacity   int
	size       int
	cache      map[int]*Node
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		head:     new(Node),
		tail:     new(Node),
		cache:    make(map[int]*Node),
	}

	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; !ok {
		return -1
	} else {
		lru.remove(node)
		lru.addToHead(node)
		return node.Val
	}
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		node.Val = value
		lru.remove(node)
		lru.addToHead(node)
	} else {
		newNode := &Node{
			Key: key,
			Val: value,
		}
		lru.cache[key] = newNode
		lru.size++
		lru.addToHead(newNode)

		if lru.size > lru.capacity {
			lastNode := lru.tail.Prev
			delete(lru.cache, lastNode.Key)
			lru.remove(lastNode)
			lru.size--
		}

	}
}

func (lru *LRUCache) remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) addToHead(node *Node) {
	node.Next = lru.head.Next
	node.Prev = lru.head

	lru.head.Next.Prev = node
	lru.head.Next = node
}
