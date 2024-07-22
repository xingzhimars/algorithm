package main

import "fmt"
import "bufio"
import "os"
import "sort"

const N = 1e5 + 10
const M = 2e5 + 10

var n, m int

// 并查集的p集合
var p [N]int

type Edge struct {
	a, b, c int
}

var Edges [M]*Edge

func kruskal() {
	sort.Slice(Edges[0:m], func(i, j int) bool {
		return Edges[i].c < Edges[j].c
	})

	for i := 0; i < N; i++ {
		p[i] = i
	}

	res := 0
	cnt := 0
	for i := 0; i < m; i++ {
		a := Edges[i].a
		b := Edges[i].b
		c := Edges[i].c

		a = find(a)
		b = find(b)

		if a != b {
			// a和b不在一个集合中
			res += c
			p[a] = b
			cnt++
		}
	}
	if cnt < n-1 {
		// 加入到集合中的点少于n-1，说明无法构造最小生成树
		fmt.Println("impossible")
	} else {
		fmt.Println(res)
	}
}

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(r, &a, &b, &c)

		Edges[i] = &Edge{
			a: a,
			b: b,
			c: c,
		}
	}

	kruskal()
}
