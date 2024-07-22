package main

import "fmt"
import "bufio"
import "os"

const N = 1e5 + 10

const M = 2e5 + 10

// 稀疏图，使用邻接表
var h [N]int
var e, ne [M]int
var idx int

var n, m int

// 染色，color[i]: 0未染色 1染白色 2染黑色
var color [N]int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func dfs(u, c int) bool {
	color[u] = c

	// 连接边也染色
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		// 如果没有染色，就将染成另外一种颜色；
		// 如果染色了，就判断颜色是否跟u一样
		if color[j] == 0 {
			// 染色值，1或者2，这里就用3减去另外一个
			if !dfs(j, 3-c) {
				return false
			}
		} else if color[j] == c {
			return false
		}

	}
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n, &m)

	for i := 0; i < N; i++ {
		h[i] = -1
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		add(a, b)
		add(b, a)
	}
	// 记录染色过程中，是否遇到问题
	flag := true
	for i := 1; i <= n; i++ {
		// 未染色
		if color[i] == 0 {
			if !dfs(i, 1) {
				flag = false
				break
			}
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
