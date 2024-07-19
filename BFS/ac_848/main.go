package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

var n, m int

var h, e, ne [N]int
var idx int

var queue, d [N]int
var hh, tt int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func topSort() bool {
	hh, tt := 0, -1

	for i := 1; i <= n; i++ {
		if d[i] == 0 {
			tt++
			queue[tt] = i
		}
	}

	for hh <= tt {
		t := queue[hh]
		hh++

		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]
			d[j]--

			if d[j] == 0 {
				tt++
				queue[tt] = j
			}
		}
	}

	// 判断是否所有节点都入队列过了
	return tt == n-1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for m > 0 {
		m--
		var a, b int
		fmt.Fscan(r, &a, &b)

		add(a, b)
		// 更新入度
		d[b]++
	}

	if topSort() {

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", queue[i])
		}

	} else {
		fmt.Println(-1)
	}

}
