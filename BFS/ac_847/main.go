package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

var n, m int

// 有向图，e,ne都是N大小
var h, e, ne [N]int
var idx int

var queue [N]int
var hh, tt int

// 距离
var d [N]int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func bfs() int {
	hh, tt := 0, -1

	tt++
	queue[tt] = 1

	d[1] = 0

	for hh <= tt {
		t := queue[hh]
		hh++

		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			if d[j] == -1 {
				d[j] = d[t] + 1

				tt++
				queue[tt] = j
			}
		}
	}
	return d[n]
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	//  h数组的初始化，必须再建边之前啊
	for i := 0; i < N; i++ {
		h[i] = -1
		d[i] = -1
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)

		add(a, b)
	}

	ans := bfs()
	fmt.Println(ans)
}
