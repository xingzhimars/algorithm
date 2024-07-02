package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

var n, m int
var p [N]int

// 集合的大小，只保证根节点是有意义的即可
// size[a]：根节点为a的集合中，连通点的个数
var size [N]int

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &n, &m)

	for i := 1; i <= n; i++ {
		p[i] = i
		size[i] = 1
	}

	for m > 0 {
		m--
		var op string
		fmt.Fscan(rr, &op)

		if op == "C" {
			var a, b int
			fmt.Fscan(rr, &a, &b)
			if find(a) == find(b) {
				continue
			}
			size[find(b)] += size[find(a)]
			p[find(a)] = find(b)

		} else if op == "Q1" {
			var a, b int
			fmt.Fscan(rr, &a, &b)
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}

		} else if op == "Q2" {
			var a int
			fmt.Fscan(rr, &a)
			fmt.Println(size[find(a)])
		}
	}
}
