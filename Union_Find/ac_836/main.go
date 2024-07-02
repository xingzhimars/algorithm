package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

var n, m int

// 每个元素的父节点是什么
var p [N]int

func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rr, &n, &m)

	// 初始化，每个数是一个集合
	for i := 1; i <= n; i++ {
		p[i] = i
	}

	for m > 0 {
		m--

		var op string
		var a, b int
		fmt.Fscan(rr, &op, &a, &b)

		if op == "M" {
			if find(a) != find(b) {
				// p[a] = b
				p[find(a)] = find(b)
			}
		} else if op == "Q" {
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}

		}

	}
}

// 核心操作
// 返回x所在集合的编号，或者说返回x的祖宗节点
// 在这个过程中加上路径压缩
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}
