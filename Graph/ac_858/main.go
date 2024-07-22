package main

import "fmt"
import "bufio"
import "os"

const N = 510
const INF = 0x3f3f3f3f

// 稠密图，使用邻接矩阵
var g [N][N]int

// 集合
var st [N]bool
var dist [N]int

var n, m int

func prim() int {
	for i := 0; i < N; i++ {
		dist[i] = INF
	}

	res := 0
	for i := 0; i < n; i++ {

		// 遍历n个点，找到距离集合最近的点
		// t表示在集合中的点，等于-1时，表示第一次，此时集合中没有点
		t := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}
		// 集合中有节点了，但是此时其他点距离集合最近的为INF，表示此时无法构造最小生成树
		if i > 0 && dist[t] == INF {
			return INF
		}

		// 将边权加起来，第一个点没有边
		if i > 0 {
			res += dist[t]
		}

		// 用t更新其他点到集合的距离
		for j := 1; j <= n; j++ {
			// 自环不要更新，否则res会变化，当然也可以将res计算提到前面去
			if t != j {
				dist[j] = min(dist[j], g[t][j])
			}
		}

		// 将t加入到集合中
		st[t] = true
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	// 邻接矩阵初始化
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = INF
		}
	}

	for m > 0 {
		m--
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)

		// 无向图是特殊的有向图，因此在赋值的时候，添加一条从a到b和一条从b到a的边
		// 本题中存在重边，因此在赋值的时候，取最小值即可
		g[a][b] = min(g[a][b], c)
		g[b][a] = g[a][b]

	}

	t := prim()

	if t == INF {
		fmt.Println("impossible")
	} else {
		fmt.Println(t)
	}

}
