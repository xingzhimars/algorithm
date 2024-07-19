package main

import "fmt"
import "bufio"
import "os"

const N = 2010
const M = 10010

var h [N]int
var e, ne, w [M]int
var idx int

var st [N]bool
var dist [N]int
var cnt [N]int

var queue []int

// var hh, tt int

var n, m int

func add(a, b, c int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	w[idx] = c
	idx++
}

func spfa() bool {
	// hh, tt := 0, -1

	for i := 1; i <= n; i++ {
		// tt++
		queue[i-1] = i
		st[i] = true
	}

	for len(queue) > 0 {
		t := queue[0]
		queue = queue[1:]

		st[t] = false

		for i := h[t]; i != -1; i = ne[i] {
			j := e[i]

			if dist[j] > dist[t]+w[i] {
				dist[j] = dist[t] + w[i]
				cnt[j] = cnt[t] + 1
				if cnt[j] >= n {
					return true
				}
				if !st[j] {
					// tt++
					// queue[tt] = j
					queue = append(queue, j)

					st[j] = true
				}
			}
		}
	}

	return false
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &n, &m)

	queue = make([]int, n+10)

	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0

	for i := 0; i < m; i++ {
		var a, b, w int
		fmt.Fscan(r, &a, &b, &w)

		add(a, b, w)
	}

	if spfa() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
