package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const N = 1e6 + 10

// 由于是稀疏图，使用邻接表
var h, e, ne, w [N]int
var idx int

var dist [N]int
var st [N]bool

var n, m int

func add(a, b, c int) {
	e[idx] = b
	ne[idx] = h[a]
	w[idx] = c
	h[a] = idx
	idx++
}

func dijkstra() int {
	dist[1] = 0
	for i := 2; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	// 堆优化
	var queue Pq
	heap.Push(&queue, &Pair{
		node:     1,
		distance: 0,
	})

	for queue.Len() > 0 {
		t := heap.Pop(&queue).(*Pair)
		if st[t.node] {
			continue
		}
		st[t.node] = true

		for i := h[t.node]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > t.distance+w[i] {
				dist[j] = t.distance + w[i]

				heap.Push(&queue, &Pair{
					node:     j,
					distance: dist[j],
				})
			}
		}
	}

	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for m > 0 {
		m--
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)
		add(a, b, c)
	}

	ans := dijkstra()
	fmt.Println(ans)
}

// golang的优先队列
type Pair struct {
	node, distance int
}

type Pq []*Pair

var _ heap.Interface = (*Pq)(nil)

func (pq Pq) Len() int {
	return len(pq)
}

func (pq Pq) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq Pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Pq) Push(x interface{}) {
	if v, ok := x.(*Pair); ok {
		*pq = append(*pq, v)
	}
}

func (pq *Pq) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}
