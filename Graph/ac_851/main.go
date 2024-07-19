package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

var h, e, ne, w [N]int
var idx int

var st [N]bool
var dist [N]int

var queue [N]int
var hh, tt int

var n, m int

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func spfa() {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0

	hh, tt := 0, -1
	tt++
	queue[tt] = 1
	st[1] = true

	for hh <= tt {
		t := queue[hh]
		hh++
		st[t] = false

		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]

			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]

				if !st[j] {
					tt++
					queue[tt] = j

					st[j] = true
				}
			}
		}
	}

	if dist[n] == 0x3f3f3f3f {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}

}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &n, &m)

	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for i := 0; i < m; i++ {
		var a, b, w int
		fmt.Fscan(r, &a, &b, &w)

		add(a, b, w)
	}

	spfa()

}
