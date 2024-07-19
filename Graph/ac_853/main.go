package main

import "fmt"
import "bufio"
import "os"

const N = 510
const M = 10010

type Edge struct {
	a, b, w int
}

var dist [N]int
var backup [N]int

var edges [M]*Edge

var n, m, k int

func bellmanFord() {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0

	for i := 0; i < k; i++ {
		// 切片才能copy
		copy(backup[0:], dist[0:])
		for j := 0; j < m; j++ {
			e := edges[j]
			// 在j遍历的过程中，dist[b]始终在变化，有可能在计算其他点时，使用了已经变化后的dist[b]
			// 这个是错误的，因此还得用上次的dist
			dist[e.b] = min(dist[e.b], backup[e.a]+e.w)
		}
	}

	// 假设dist[5]为0x3f3f3f3f，节点5到节点n的权重为-2，那么dist[n]就是0x3f3f3f3f - 2，实际上也是无穷大，
	// 但是却没法等于0x3f3f3f3f
	if dist[n] > 0x3f3f3f3f/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m, &k)

	for i := 0; i < m; i++ {
		var a, b, w int
		fmt.Fscan(r, &a, &b, &w)

		edges[i] = &Edge{
			a: a,
			b: b,
			w: w,
		}
	}

	bellmanFord()
}
