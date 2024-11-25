package main

import "fmt"

func main() {
	nums := []int{3, 0, 1, 0}
	k := 1

	res := topKFrequent(nums, k)
	fmt.Println(res)
}

type Node struct {
	Val  int
	Freq int
}

var heap []*Node
var size int

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		v, ok := mp[nums[i]]
		if !ok {
			mp[nums[i]] = 1
		} else {
			mp[nums[i]] = v + 1
		}
	}

	nodes := make([]*Node, 0)

	for k, v := range mp {
		node := &Node{
			Val:  k,
			Freq: v,
		}
		nodes = append(nodes, node)
	}

	n := len(nodes)

	heap = make([]*Node, k+10)
	size = k

	for i := 0; i < n; i++ {
		if i < k {
			heap[i+1] = nodes[i]
		} else {
			buildHeap()
			if heap[1].Freq > nodes[i].Freq {
				continue
			}
			heap[1] = nodes[i]
			down(1)
		}
	}

	ans := make([]int, k)
	for i := 1; i <= k; i++ {
		ans[i-1] = heap[i].Val
	}
	return ans
}

// 建立
func buildHeap() {
	for i := size / 2; i > 0; i-- {
		down(i)
	}
}

// 调整
func down(u int) {
	t := u
	if u*2 <= size && heap[u*2].Freq < heap[t].Freq {
		t = u * 2
	}
	if u*2+1 <= size && heap[u*2+1].Freq < heap[t].Freq {
		t = u*2 + 1
	}
	if u != t {
		heap[u], heap[t] = heap[t], heap[u]
		down(t)
	}
}
