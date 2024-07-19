package main

import "fmt"
import "bufio"
import "os"

const N = 510

var g [N][N]int

var st [N]bool

var dist [N]int

var n, m int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dijkstra() int {
	dist[1] = 0

	for i := 2; i <= n; i++ {
		dist[i] = 0x3f3f3f3f
	}

	// 遍历n次
	for i := 0; i < n; i++ {

		// 1. 找出不存在集合s中，且距离最近的点
		t := -1
		for j := 1; j <= n; j++ {
			// t == -1表示首次
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		st[t] = true

		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+g[t][j])
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

	// 初始化g
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == j {
				// 自环
				g[i][j] = 0
			} else {
				g[i][j] = 0x3f3f3f3f
			}
		}
	}

	for m > 0 {
		m--
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)

		// 题目说存在重边，那么取最小边即可
		g[a][b] = min(g[a][b], c)
	}

	ans := dijkstra()
	fmt.Println(ans)
}
